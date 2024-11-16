package main

import (
	"fmt"
	"io"
	"os"
)

// go run main.go myfile.txt

func main() {
	fileName := os.Args[1]
	file := readFromFile(fileName)
	io.Copy(os.Stdout, file)
}

func readFromFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		//Do something
		fmt.Println(err)
		os.Exit(1)
	}
	return file
}
