package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

//--|receiver|-|func|--|arguments-passed|---
func (p Point) move(first int, second int) {
	p.x += first
	p.y += second
}

//---|receiver|--|func|--|arguments-passed|---
func (p *Point) movePtr(first int, second int) {
	p.x += first
	p.y += second
}

func main() {
	p1 := Point{1, 2}
	fmt.Printf("%+v \n", p1)

	p1.move(2, 3)
	fmt.Printf("%+v \n", p1)

	p1.movePtr(2, 3)
	fmt.Printf("%+v \n", p1)
}
