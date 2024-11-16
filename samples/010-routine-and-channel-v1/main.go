package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://leetcode.com",
		"https://cert.pl",
		"https://nasdaqomxnordic.com",
		"https://go.dev",
		"https://github.com",
	}

	// channel string
	c := make(chan string)

	for _, link := range links {
		go checkStatus(link, c)
	}

	for i := 0; i < len(links); i++ { // wait for 5 messages
		fmt.Println(<-c) // wait for message (blocks)
	}
}

func checkStatus(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		c <- "some error" // send message
		return
	}
	fmt.Println("Status for", link, ":", resp.StatusCode)
	c <- "ok" // send message
}
