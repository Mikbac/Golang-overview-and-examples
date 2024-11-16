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

	for l := range c { // infinitive loop, better than 011-routine-and-channel-v2
		go checkStatus(l, c)
	}
}

func checkStatus(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		c <- link // send message
		return
	}
	fmt.Println("Status for", link, ":", resp.StatusCode)
	c <- link // send message
}
