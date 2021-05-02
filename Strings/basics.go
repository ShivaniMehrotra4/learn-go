package main

import (
	"fmt"
)

func main() {
	a_string := "Hello, I am new to Go"
	fmt.Println(a_string)
	fmt.Println(len(a_string))
	fmt.Printf("a_string[0] = %c of type = %T \n", a_string[0], a_string[0])

	// Uncomment below line to understand the error.
	// STRINGS IN GO ARE IMMUTABLE
	// a_string[5] = "b"
	fmt.Println(a_string + "16")

	//Slices
	fmt.Println(a_string[4:10])
	fmt.Println(a_string[:6])
	fmt.Println(a_string[6:])

	//Multi-line-Strings
	say := `
	I
	Am
	New
	To
	Go
	...`

	fmt.Println(say)

}
