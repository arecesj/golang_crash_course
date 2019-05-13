package main

import "fmt"

func main() {
	name := "Juan"
	// size := 1.45
	var size float32 = 1.45
	var age int32 = 28
	var isLame = true
	isLame = false
	address, email := "SF", "123@funst.com"

	fmt.Println(name, age, isLame, address, email)
	fmt.Printf("%T\n", size)
}
