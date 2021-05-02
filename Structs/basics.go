// Structs: combining several types (fields) into one data type

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

func main() {
	s1 := Stock{"Reliance", 45.8, "$", true}
	fmt.Println(s1)
	fmt.Printf("%+v \n", s1)
	fmt.Println(s1.money)

	s2 := Stock{
		name:     "Airtel",
		money:    33.6,
		currency: "P",
		lena_hai: false,
	}
	fmt.Println(s2)
	fmt.Printf("%+v \n", s2)

	s3 := Stock{}
	fmt.Println(s3)
	fmt.Printf("%+v \n", s3)

}
