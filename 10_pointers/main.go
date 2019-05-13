package main

import "fmt"

func main() {
	// Pointer allows you to point to the memory address / location of a value
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read value from address
	fmt.Println(*b)

	// Change val with pointer
	*b = 10

	fmt.Println(a)
}
