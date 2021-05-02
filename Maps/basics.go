package main

import "fmt"

func main() {
	element := map[string]int{
		"One":   1,
		"Two":   2,
		"three": 3,
	}

	fmt.Println(element)

	fmt.Println(len(element))

	fmt.Println(element["Two"])

	fmt.Println(element["Four"])

	value, status := element["One"]
	if !status {
		fmt.Println("Not present")
	} else {
		fmt.Println(value)
	}

	element["Two"] = 22
	element["Four"] = 4
	fmt.Println(element)

	for i := range element {
		fmt.Println(i)
	}

	for i, j := range element {
		fmt.Println(i, j)
	}

	delete(element, "One")
	fmt.Println(element)
}
