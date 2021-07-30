package main

func main() {
	average := average(12, 13, 14, 15, 16, 17, 19)
	println(int(average))
}
func average(numbers ...int) float64 {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return float64(sum) / float64(len(numbers))
}
