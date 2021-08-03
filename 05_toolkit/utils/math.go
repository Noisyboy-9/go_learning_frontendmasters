package utils

// Sum calculates the sum of two or more numbers
func Sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

// Subtract calculates the subtraction of two numbers
func Subtract(num1, num2 int) int {
	return num1 - num2
}

// Multiply calculates the multiplication of two or more numbers
func Multiply(nums ...int) int {
	result := 1

	for _, num := range nums {
		result *= num
	}

	return result
}

// Divide calculates the division of two numbers
func Divide(num1, num2 int) float64 {
	return float64(num1) / float64(num2)
}
