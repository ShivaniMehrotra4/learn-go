package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 353
	}()

	receive := <-ch // We are trying to receive from the channel here.
	fmt.Printf("Here, I got %v from channel \n ", receive)
}
