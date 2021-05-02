package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://facebook.com"
	resp, error := contentType(url)
	if error != nil {
		fmt.Printf("Error is : %s \n", error)
	} else {
		fmt.Printf("Response is : %s \n", resp)
	}
}

func contentType(url string) (string, error) {
	response, error := http.Get(url)
	if error != nil {
		return "", error
	}
	defer response.Body.Close()
	content_type := response.Header.Get("Content-Type")
	if content_type == "" {
		return "", fmt.Errorf("Can't get the content-type header")
	}
	return content_type, nil
}

// knoldus@shivani-mehrotra:~/Documents/PlayGround/learn-go/Functions $ go run challenge.go
// Error is : Get "https:///golang.org": http: no Host in request URL
// knoldus@shivani-mehrotra:~/Documents/PlayGround/learn-go/Functions $ go run challenge.go
// Response is : text/html; charset=utf-8
// knoldus@shivani-mehrotra:~/Documents/PlayGround/learn-go/Functions $ go run challenge.go
// Response is : text/html; charset=ISO-8859-1
// knoldus@shivani-mehrotra:~/Documents/PlayGround/learn-go/Functions $ go run challenge.go
// Response is : text/html; charset="utf-8"
