//  Channels

// Two types:
// 1. buffered
// 2. unbuffered

// -> When we push a value across unbuffered channel, we get BLOCKED until someone receives that input on the other side.

// -> Each Buffered channel has it's own capacity.
//    If we send value across buffered channel, we won't get blocked untilthe channel's capacity if FULL.
//    When capacity is full, ascent to a channel will be blocked.

//    Buffered channels are somewhat like BONDED QUEUES.

// We make a channel with the "make" command.

package main

import "fmt"

func main() {
	ch := make(chan int)
	<-ch
	fmt.Println("Here")
}
