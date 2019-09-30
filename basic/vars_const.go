package main

import "fmt"

func main() {
	const pi float64 = (22 / 7)
	var radius float64 = 7

	fmt.Println("Its a constant Pi ", pi)
	fmt.Println("Its a variable Radius ", radius)

	fmt.Println("So the area is ", pi*radius*radius)
	fmt.Println("And the Circumpherence is ", pi*2*radius)
}
