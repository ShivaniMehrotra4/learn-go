package main

import (
	"fmt"
)

type Stock struct {
	name     string
	money    float64
	currency string
	lena_hai bool
}

func (s *Stock) doublePrice() float64 {
	if s.lena_hai == false {
		s.money *= 1
	}
	return s.money * 2
}

func main() {
	s1 := Stock{"name", 1.0, "$", true}
	fmt.Println(s1)
	s2 := s1.doublePrice()
	fmt.Println(s2)
}
