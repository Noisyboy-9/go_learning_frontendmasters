package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

func NewUser(ID int, firstName string, lastName string, email string) *User {
	return &User{ID: ID, FirstName: firstName, LastName: lastName, Email: email}
}

type Group struct {
	role           string
	capacity       int
	users          []*User
	newestUser     *User
	spaceAvailable bool
}

func NewGroup(role string, capacity int) *Group {
	return &Group{
		role:           role,
		capacity:       capacity,
		spaceAvailable: true,
		newestUser:     nil,
	}
}

func (group *Group) addUser(user *User) error {
	if group.spaceAvailable {
		group.users = append(group.users, user)
		group.newestUser = user
	} else {
		return errors.New("the group doesn't have any space")
	}

	if len(group.users) >= group.capacity {
		group.spaceAvailable = false
	}

	return nil
}

func (group *Group) describe() string {
	return fmt.Sprintf(
		"This user group has %d users. The neweset user is %s %s. Accepting new User: %t",
		len(group.users),
		group.newestUser.FirstName,
		group.newestUser.LastName,
		group.spaceAvailable,
	)
}

func (user *User) describe() string {
	return fmt.Sprintf(
		"ID:%d\nName: %s %s\nEmail: %s",
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
	)
}

func main() {
	group := NewGroup("admin", 5)

	user1 := NewUser(1, "sina", "shariati", "sina.shariati@yahoo.com")
	user2 := NewUser(2, "adib", "sobhani", "adib.sobhani@yahoo.com")
	user3 := NewUser(3, "aryan", "sohrabi", "aryan.sohrabi@yahoo.com")
	user4 := NewUser(4, "john", "doe", "john.doe@yahoo.com")
	user5 := NewUser(5, "jane", "doe", "jane.doe@yahoo.com")

	if err := group.addUser(user1); err != nil {
		err := recover()
		panic(err)
	}

	if err := group.addUser(user2); err != nil {
		err := recover()
		panic(err)
	}

	if err := group.addUser(user3); err != nil {
		err := recover()
		panic(err)
	}

	if err := group.addUser(user4); err != nil {
		err := recover()
		panic(err)
	}

	if err := group.addUser(user5); err != nil {
		err := recover()
		panic(err)
	}

	fmt.Println(group.describe())
	fmt.Println(user1.describe())
}
