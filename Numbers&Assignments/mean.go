package main // means stand-alone executable not as a shared library.

import (
	"fmt"
)

func main() {
	var x int
	var y int

	x = 5
	y = 7

	fmt.Printf("x=%v , type of %T\n", x, x)
	fmt.Printf("y=%v , type of %T\n", y, y)

	var mean int
	mean = (x + y) / 2
	fmt.Printf("Mean of %v and %v = %v , type of %T \n", x, y, mean, mean)

}
