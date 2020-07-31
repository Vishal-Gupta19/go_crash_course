package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// Define map
	emails := make(map[string]string)

	// Assign
	emails["Vishal"] = "vishal19@gmail.com"
	emails["Tejas"] = "tejas05@gmail.com"

	fmt.Println(emails)

	delete(emails, "Vishal")
	fmt.Println(emails)

}
