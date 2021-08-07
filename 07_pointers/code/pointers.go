package main

import "fmt"

//func main() {
//	name := "sina shariati"
//	namePointer := &name
//
//	nameValue := *namePointer
//}

// // ******************************************************

// func main() {
// 	var name string = "Beyonce"
// 	var namePointer *string = &name
// 	var nameValue = *namePointer

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name Value:", nameValue)

// }

// // ******************************************************

//func changeName(stringPointer *string) {
//	*stringPointer = strings.ToUpper(*stringPointer)
//}
//
//func main() {
//	name := "Elvis"
//	changeName(&name)
//	fmt.Println(name)
//}

// Coordinates model a place in earth.
type Coordinates struct {
	X, Y float64
}

func mirrorCoordinate(coordinatePointer *Coordinates) *Coordinates {
	temp := coordinatePointer.X
	(*coordinatePointer).X = (*coordinatePointer).Y
	(*coordinatePointer).Y = temp
	return coordinatePointer
}

func NewCoordinates(x float64, y float64) *Coordinates {
	return &Coordinates{X: x, Y: y}
}

func main() {
	myHomePointer := NewCoordinates(12, 20)
	mirrorCoordinate(myHomePointer)
	fmt.Print(*myHomePointer)
}
