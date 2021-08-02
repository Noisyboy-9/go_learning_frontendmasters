package main

func main() {
	average := average(12, 13, 14, 15, 16, 17, 19)
	println(average)
}
func average(numbers ...float64) float64 {
	sum := 0.0

	for _, number := range numbers {
		sum += number
	}

	return sum / float64(len(numbers))
}
