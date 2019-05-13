package main

import (
	"fmt"
	"strconv"
)

// Define person struct
// type Person struct {
// 	firstName string
// 	lastName  string
// 	city      string
// 	gender    string
// 	age       int
// }
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver) --> takes values it receives and creates an output. No updating. "Make calculations"
func (p Person) greet() string {
	// Imported a package to convert age to a string bc you cannot have two diff types concat
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old"
}

// Has birthday method (pointer receiver) --> because we are changing something within the struct
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}

	// Alternative
	// remember you can have vars in there
	// person1 := Person{"Samantha", "Smith", "Boston", "f", 25}
	person2 := Person{"Bob", "Johnson", "New York", "m", 30}

	fmt.Println(person1)

	// Get single field
	fmt.Println(person1.firstName)
	// Update a field
	// person1.age++
	fmt.Println(person1)
	person1.hasBirthday()
	person1.getMarried("Williams")
	person2.getMarried("Thompson")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
