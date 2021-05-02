package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p *Point) MovePOint(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

type Square struct {
	centre Point
	length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if x < 0 {
		return nil, fmt.Errorf("%v cannot be less than 0", x)
	}
	if y < 0 {
		return nil, fmt.Errorf("%v cannot be less than 0", y)
	}
	if length < 0 {
		return nil, fmt.Errorf("%v cannot be less than 0", length)
	}

	sq := &Square{
		centre: Point{x, y},
		length: length,
	}
	return sq, nil
}

func (s *Square) Move(dx int, dy int) {
	s.centre.MovePOint(dx, dy)
}

func (s *Square) Area() int {
	return s.length * s.length
}

func main() {
	s1, error := NewSquare(1, 2, 2)
	if error != nil {
		fmt.Printf("Error is %s \n", error)
	}

	fmt.Println(s1)

	// MOve Square
	s1.Move(3, 6)
	fmt.Println("Calculating Area")

	fmt.Println()
	fmt.Printf("Area of Square is: %v", s1.Area())

}
