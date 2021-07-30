package main

import "fmt"

func main() {
	if err := someFunction(); err != nil {
		fmt.Println("something bad")
	}
}
