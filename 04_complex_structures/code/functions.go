package main

import "fmt"

func main() {
	average := average(19, 20, 21, 23, 24, 26)
	fmt.Print(average)
}

func average(numbers ...int) float64 {
	sum := 0

	for _, age := range numbers {
		sum += age
	}

	return float64(sum) / float64(len(numbers))
}
