package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " age " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouse string) {
	if p.gender == "F" {
		return
	} else {
		p.lastName = spouse
	}
}

func main() {
	person1 := Person{firstName: "Vishal", lastName: "Gupta", city: "Bangalore", gender: "M", age: 24}

	fmt.Println(person1)
	person1.hasBirthday()
	person1.getMarried("Hah")
	fmt.Println(person1.greet())
}
