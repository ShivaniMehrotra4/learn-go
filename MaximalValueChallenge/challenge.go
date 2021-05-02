package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	max := nums[0]
	for index := range nums[1:] {
		if max < nums[index] {
			max = nums[index]
		}
	}
	fmt.Println(max)

	//OR

	max1 := nums[0]
	for _, value := range nums[1:] {
		if max1 < value {
			max1 = value
		}
	}
	fmt.Println(max1)
}
