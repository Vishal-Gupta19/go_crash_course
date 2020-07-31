package main

import "fmt"

func main() {
	//Arrays
	var fruitArray [2]string

	// Assign values
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	//vegArray := []string{"Carrot", "beans"}
	fmt.Println("Hello World")
	fmt.Println(fruitArray)

	//Slices
	vegArray := []string{"Carrot", "beans"}
	fmt.Println(vegArray)

}
