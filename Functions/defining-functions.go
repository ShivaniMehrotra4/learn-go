package main

import (
	"fmt"
)

func main() {
	a, b := 3, 4
	sum := add(a, b)
	fmt.Println(sum)

	p, q := 10, 2
	modulus, division := return_multiple(p, q) // PANIC ERROR : divide by zero error
	fmt.Println(modulus, division)
}

func add(a int, b int) int {
	return a + b
}

func return_multiple(a int, b int) (int, int) {
	return a % b, a / b
}
