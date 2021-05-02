// COncurrency is the composition of independently executed processes.

// Parrallelism is the simultaneous execution of processes.

package main

import (
	"fmt"
	"net/http"
	"sync" //NEW
)

func main() {
	urls := []string{"https://linkedin.com", "https://facebook.com", "https://google.com"}

	//NEW
	var wait_group sync.WaitGroup //NEW
	for _, url := range urls {
		wait_group.Add(1) // NEW
		go func(url string) {
			returnType(url)
			wait_group.Done() // NEW
		}(url)
	}
	wait_group.Wait() // NEW
}

func returnType(url string) {
	response, error := http.Get(url)
	if error != nil {
		fmt.Printf("error is %s \n", error)
	}

	defer response.Body.Close()
	content_type := response.Header.Get("content-type")
	fmt.Printf("%s -> %s \n", url, content_type)
}

// We use time command to run this.

/*
knoldus@shivani-mehrotra:~/Documents/PlayGround/learn-go/Concurrency $ time go run goroutines.go

real    0m0.366s
user    0m0.491s
sys     0m0.077s
*/

// Nothing is printed out here, THE GO RUNTIME WILL NOT WAIT FOR THE GO ROUTINES TO END.
// we can use sleep to wait for them, but it's a bad practice of synchronization.
// Next, in modified-goroutines.go, we use sync package, but still the best recommended way is to use channels.

// Also, we use WAITGROUP in order to wait for go routines.
// -----------------------------------------------------------------------------------------------
