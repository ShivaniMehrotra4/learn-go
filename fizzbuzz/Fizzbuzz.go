package main

import (
	"fmt"
)

func main() {
	num1 := 1
	for num1 <= 20 {
		if num1%3 == 0 && num1%5 == 0 {
			fmt.Println("fizz buzz")
		} else if num1%5 == 0 {
			fmt.Println("buzz")
		} else if num1%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(num1)
		}
		num1++
	}
}
