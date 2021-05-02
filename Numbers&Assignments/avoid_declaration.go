package main

import (
	"fmt"
)

func main() {

	// avoiding declaration and then assignment to just assignment.
	x := 5.0
	y := 7.0

	fmt.Printf("x=%v , type of %T\n", x, x)
	fmt.Printf("y=%v , type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("Mean of %v and %v = %v , type of %T \n", x, y, mean, mean)

	// decalaring in single line
	p, q := 3, 7
	fmt.Println(p + q)
}
