package main

import (
	"fmt"
)

func main() {
	for x := 0; x < 5; x++ {
		fmt.Println(x)
	}

	print("----------------")
	for x := 0; x < 5; x++ {
		if x > 3 {
			break
		}
		fmt.Println(x)
	}

	print("----------------")
	for x := 0; x < 5; x++ {
		if x > 1 {
			continue
		}
		fmt.Println(x)
	}

	print("----------------")
	a := 7
	for a > 2 {
		fmt.Println(a)
		a--
	}

	print("----------------")
	b := 1
	for {
		fmt.Println(b)
		if b != 3 {
			break
		}
		b++
	}
}
