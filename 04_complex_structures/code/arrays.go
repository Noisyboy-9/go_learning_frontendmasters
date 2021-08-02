package main

import (
	"fmt"
)

func main() {
	scores := []int{9, 2, 7, 8, 10}

	for _, value := range scores {
		fmt.Println(value)
	}
}
