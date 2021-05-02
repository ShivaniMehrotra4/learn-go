package main

import (
	"fmt"
)

// Defer: close all unnecessary resources : not just memory; but like files,etc.

func main() {
	worker()
	fmt.Println("Inside main")
}

func worker() {
	defer cleanup("A")
	defer cleanup("B")
	fmt.Println("Inside worker")
}

func cleanup(str string) {
	fmt.Printf("Inside cleanup : Cleaning up %s \n", str)
}
