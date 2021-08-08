package main

import "fmt"

type User struct {
	name  string
	age   int
	email string
}

func NewUser(name string, age int, email string) *User {
	return &User{name: name, age: age, email: email}
}

func (user *User) sayHello() {
	fmt.Printf("%s says hello", user.name)
}

func main() {
	user := NewUser("sina shariati", 19, "sina.shariati@yahoo.com")
	user.sayHello()
}
