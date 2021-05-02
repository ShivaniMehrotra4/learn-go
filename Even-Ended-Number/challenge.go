package main

import (
	"fmt"
)

func main() {
	var count int
	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			product := i * j
			str_product := fmt.Sprintf("%d", product)
			if str_product[0] == str_product[len(str_product)-1] {
				count++
			}
		}
	}
	fmt.Printf("The result is: %d", count)
}
