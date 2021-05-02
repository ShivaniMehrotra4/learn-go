package main

import (
	"fmt"
)

func main() {
	x := 10

	if x > 9 {
		fmt.Println("x is big")
	}

	y := 0
	if x < 100 && y != 0 {
		fmt.Print("hello, they satisfy the conditions.\n")
	} else if x == 10 || y >= -3 {
		fmt.Print("hello, they still satisfy the conditions.\n")
	} else {
		fmt.Print("hello, they do not satisfy the conditions.\n")
	}

}
