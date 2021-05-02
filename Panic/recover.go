package main

import (
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(list)

	v := safeType(list, 8)
	fmt.Println(v)

}

func safeType(values []int, index int) int {

	// Create an ANONYMOUS FUNCTION
	defer func() {
		if error := recover(); error != nil {
			fmt.Printf("Error is : %s \n ", error)
		}
	}()
	return values[index]
}
