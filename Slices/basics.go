package main

import (
	"fmt"
)

/*
Slices are sequence of items.
All items in a slice have to be of the same type.
*/
func main() {
	// Creating Slices
	str := []string{"I", "AM", "NEW", "TO", "GO"}
	fmt.Println(str)
	fmt.Println(str[2:])
	// fmt.Println(str[5]) // results in PANIC
	fmt.Println(str[:3])
	fmt.Println(len(str))

}
