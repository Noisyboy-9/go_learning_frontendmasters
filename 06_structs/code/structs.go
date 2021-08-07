package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// Group is for grouping the users with the same role
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func NewGroup(role string, users []User, newestUser User, spaceAvailable bool) *Group {
	return &Group{role: role, users: users, newestUser: newestUser, spaceAvailable: spaceAvailable}
}

func describeUser(user User) string {
	return fmt.Sprintf(
		"ID:%d\nName: %s %s\nEmail: %s",
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
	)
}

func describeGroup(group Group) string {
	if len(group.users) >= 2 {
		group.spaceAvailable = true
	}

	return fmt.Sprintf(
		"This user group has %d users. The neweset user is %s %s. Accepting new User: %t",
		len(group.users),
		group.newestUser.FirstName,
		group.newestUser.LastName,
		group.spaceAvailable,
	)
}

func main() {
	user1 := User{
		ID:        1,
		FirstName: "sina",
		LastName:  "shariati",
		Email:     "sina.shariati@yahoo.com",
	}

	user2 := User{
		ID:        2,
		FirstName: "adib",
		LastName:  "sobhani",
		Email:     "adib.sobhani@yahoo.com",
	}

	user3 := User{
		ID:        3,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@yahoo.com",
	}

	user4 := User{
		ID:        4,
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane.doe@yahoo.com",
	}

	group1 := Group{
		role:           "admin",
		users:          []User{user2, user1},
		newestUser:     user1,
		spaceAvailable: false,
	}
	group2 := Group{
		role:           "normal-user",
		users:          []User{user3, user4},
		newestUser:     user4,
		spaceAvailable: true,
	}

	println(describeGroup(group1))
	println(describeGroup(group2))

	println("---------")
	println(describeGroup(group2))
}
