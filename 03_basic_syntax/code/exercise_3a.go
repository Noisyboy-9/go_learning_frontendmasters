package main

func isOdd(number int) bool {
	return number%2 == 1
}

func main() {
	mySentence := "This is my sentence what do you think?"

	for index, letter := range mySentence {
		if isOdd(index) {
			println(string(letter))
		}
	}
}
