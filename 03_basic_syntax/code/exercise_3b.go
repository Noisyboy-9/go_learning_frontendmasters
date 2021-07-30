package main

import "fmt"

func main() {
	var name string
	var city string
	var cityLivingDuration int
	var weatherIdea bool

	fmt.Print("please input your name:\t")
	fmt.Scan(&name)

	fmt.Print("please input your city:\t")
	fmt.Scan(&city)

	fmt.Print("How many years have you been living in ", city, ":\t")
	fmt.Scan(&cityLivingDuration)

	fmt.Print("is the weather good there:\t")
	fmt.Scan(&weatherIdea)

	fmt.Printf(
		"Hi! My name is %s. I have lived in %s for %d years. They say the weather is amazing, which is %t.\n",
		name,
		city,
		cityLivingDuration,
		weatherIdea,
	)
}
