package main

func main() {
	average := threeNumberAverage(18, 19, 20)
	println(int(average))
}

func threeNumberAverage(number1, number2, number3 float64) float64 {
	sum := number1 + number2 + number3
	return sum / 3
}
