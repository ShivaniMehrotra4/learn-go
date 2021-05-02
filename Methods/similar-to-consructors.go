package main

import (
	"fmt"
	"os"
)

type Point struct {
	length  int
	breadth int
}

func NewPoint(first int, second int) (*Point, error) {
	if first < 0 {
		return nil, fmt.Errorf("%v cannot be less than zero.", first)
	}
	if second < 0 {
		return nil, fmt.Errorf("%v cannot be less than zero.", second)
	}

	point := &Point{
		length:  first,
		breadth: second,
	}
	return point, nil
}

func main() {
	p, err := NewPoint(3, 8)
	if err != nil {
		fmt.Printf("ERROR: %s \n", err)
		os.Exit(1)
	}

	fmt.Println(p)
}
