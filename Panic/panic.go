package main

import (
	"fmt"
	"os"
)

func main() {
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(list)

	// fmt.Println(list[8]) // reults in PANIC: index out of range

	file, error := os.Open("no-such-file") // PANIC
	// Change fiel name to panic.go : IT WORKS...
	if error != nil {
		panic(error)
	}
	defer file.Close()
	fmt.Println("File is opened")
}
