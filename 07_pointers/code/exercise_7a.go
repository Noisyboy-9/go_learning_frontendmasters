package main

import "fmt"

// User for a client using the software
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

func updateMail(user *User) {
	user.Email = "newEmail@yahoo.com"
}

func main() {
	user1 := User{
		ID:        1,
		FirstName: "sina",
		LastName:  "shariati",
		Email:     "sina.shariati@yahoo.com",
	}

	fmt.Println(user1)
	updateMail(&user1)
	fmt.Println(user1)
}
