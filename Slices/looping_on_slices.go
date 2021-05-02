package main

import (
	"fmt"
)

func main() {
	str := []string{"I", "AM", "NEW", "TO", "GO"}
	fmt.Println(str)

	// for loop
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	// Single Range (only index)
	fmt.Println("----------------------")
	for i := range str {
		fmt.Println(i)
	}

	// Double Range
	fmt.Println("----------------------")
	for index, value := range str {
		fmt.Printf("Value at index: %d is %v \n", index, value)
	}

	// Double Range | Ignoring index
	fmt.Println("----------------------")
	for _, value := range str {
		fmt.Printf("Value is %v \n", value)
	}

	// Append element in a slice
	fmt.Println("----------------------")
	str = append(str, "Period")
	fmt.Println(str)
}
