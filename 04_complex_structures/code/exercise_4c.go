package main

import "fmt"

func main() {
	//	part 1
	scores := [5]float64{16.2, 18.69, 20, 17, 17.75}
	fmt.Println(average(scores))

	//  part 2
	animalTypeToNameMap := map[string]string{
		"doberman":     "dog",
		"goldfish":     "fish",
		"maximilianus": "dog",
	}

	fmt.Println(mapKeyExists(animalTypeToNameMap, "hama"))

	//	part 3
	mySlice := []string{"banana", "apple", "lemons", "chicken", "fruits"}

	myNewSlice := addItemToSlice(mySlice, "new item 1", "new item 2", "and even new item 3")
	fmt.Println(myNewSlice)
}

func addItemToSlice(slice []string, items ...string) []string {
	return append(slice, items...)
}

func mapKeyExists(targetMap map[string]string, key string) bool {
	_, exists := targetMap[key]
	return exists
}

func average(scores [5]float64) float64 {
	sum := 0.0

	for _, score := range scores {
		sum += score
	}

	return sum / float64(len(scores))
}
