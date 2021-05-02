// COncurrency is the composition of independently executed processes.

// Parrallelism is the simultaneous execution of processes.

package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{"https://linkedin.com", "https://facebook.com", "https://google.com"}

	for _, url := range urls {
		returnType(url)
	}
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
knoldus@shivani-mehrotra:~/Documents/PlayGround/learn-go/Concurrency $ time go run not-a-goroutine-here.go
https://linkedin.com -> text/html; charset=utf-8
https://facebook.com -> text/html; charset="utf-8"
https://google.com -> text/html; charset=ISO-8859-1

real    0m2.510s
user    0m0.552s
sys     0m0.132s
*/

// -----------------------------------------------------------------------------------------------
