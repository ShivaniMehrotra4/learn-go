// Select build-in function allows us to work with several channels together.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		ch1 <- 100

	}()
	select {
	case num := <-ch1:
		fmt.Printf("Received value %v from channel ch1 \n", num)
	case num := <-ch2:
		fmt.Printf("Received value %v from channel ch2 \n", num)

	}

	// Most common use-case of select is: Implement timeouts
	fmt.Println("___________________________________________________________________")

	out_channel := make(chan float64)
	go func() {
		time.Sleep(100 * time.Millisecond)
		out_channel <- 1.234
	}()

	select {
	case value := <-out_channel:
		fmt.Printf("Received value %v from channel ch1 \n", value)
	// case <-time.After(110 * time.Millisecond):
	// 	fmt.Println("No timeout!!!!")
	case <-time.After(11 * time.Millisecond):
		fmt.Println("timeout!!!!")
	}

}
