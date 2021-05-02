package main

import (
	"fmt"
)

func main() {
	x := 3

	switch x {
	case 1:
		fmt.Println("I")
	case 2:
		fmt.Println("AM")
	case 3:
		fmt.Println("Tie")
	default:
		fmt.Println("Default")
	}

	switch {
	case x < 10:
		fmt.Println("x is smaller than 10")
	case x > 10:
		fmt.Println("x is greater than 10")
	default:
		fmt.Println("x is equal to 10")
	}
}
