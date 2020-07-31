package main

import "fmt"

func main() {
	//Using var
	// var name = "Vishal"
	var age = 24

	// Shorthand
	name, email := "Gupta", "vishalgupta9596@gmail.com"
	size := 1.3

	fmt.Println(name, age, size, email)
	fmt.Printf("%T, %T, %T\n", name, age, size)
}
