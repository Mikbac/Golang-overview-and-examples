package main

import (
	"fmt"
	"os"
)

// go run main.go Bob

func main() {
	var arg string = os.Args[1]
	fmt.Printf("Hello %s!", arg)
}
