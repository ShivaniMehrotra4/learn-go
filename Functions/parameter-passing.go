package main

import (
	"fmt"
)

func DOdoubleATPosition(slice []int, val int) {
	slice[val] *= 2
	fmt.Println(slice)
}

func double(num int) {
	num *= 2
}

func doublePointer(num *int) {
	*num *= 2
}

func main() {
	list := []int{1, 2, 3, 5}
	DOdoubleATPosition(list, 2)
	fmt.Println(list)

	number := 100
	double(number)
	fmt.Println(number)

	doublePointer(&number)
	fmt.Println(number)
}
