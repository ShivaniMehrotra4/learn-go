package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 4; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 1; i <= 4; i++ {
		receive := <-ch // We are trying to receive from the channel here.
		fmt.Printf("Here, I got %v from channel \n ", receive)

	}

	// Another good practice: Close the channel, if we don't know how many items are there going to be.
	// The sender can signal this, by closing the channel after sending.

	closeChannel()

}

func closeChannel() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 4; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("Here, I got %v from channel \n ", i)

	}

}
