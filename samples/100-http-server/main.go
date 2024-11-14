package main

import (
	"fmt"
	"net/http"
)

// curl localhost:8080

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Golang!")
}
