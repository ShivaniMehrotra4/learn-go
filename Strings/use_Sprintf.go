package main

import (
	"fmt"
)

func main() {
	num := 16
	string_num := fmt.Sprintf("%d", num)
	fmt.Printf("value of num = %v and it's type is = %T \n", num, num)
	fmt.Printf("String Value of num = %s and it's type = %T \n", string_num, string_num)
	fmt.Println("For more clarity.......... ")
	fmt.Printf("String Value of num = %q and it's type = %T \n", string_num, string_num)

}
