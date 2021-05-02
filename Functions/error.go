package main

import (
	"fmt"
	"math"
)

func squareroot(num float64) (float64, error) {
	if num < 0 {
		return 0.0, fmt.Errorf("Square root of a negative number %v is not possible", num)
	} else {
		return math.Sqrt(num), nil
	}
}

func main() {
	num1, num2 := 4.0, -4.0

	sqrt1, error1 := squareroot(num1)
	if error1 != nil {
		fmt.Printf("Error is: %s \n", error1)
	} else {
		fmt.Printf("Square root of %f is: %f \n", num1, sqrt1)
	}

	sqrt2, error2 := squareroot(num2)
	if error2 != nil {
		fmt.Printf("Error is: %s \n", error2)
	} else {
		fmt.Printf("Square root of %f is: %f \n", num2, sqrt2)
	}
}
