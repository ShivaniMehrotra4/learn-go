package main

import (
	"fmt"
)

func main() {
	var x float64
	var y int

	x = 5.0
	y = 7

	fmt.Printf("x=%v , type of %T\n", x, x)
	fmt.Printf("y=%v , type of %T\n", y, y)

	var mean float64
	mean = (x + y) / 2.0
	fmt.Printf("Mean of %v and %v = %v , type of %T \n", x, y, mean, mean)

}
