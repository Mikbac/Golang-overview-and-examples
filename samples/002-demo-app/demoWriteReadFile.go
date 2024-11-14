package main

import (
	"fmt"
	"os"
	"strings"
)

func runWriteReadFileDemo() {
	filename := "samples/002-demo-app/resources/file1.test"
	writeToFile(filename)
	readFromFile(filename)
}

func writeToFile(filename string) {
	d1 := []byte("aaaa bbbb cccc\ngo\n444 555 777")
	err := os.WriteFile(filename, d1, 0666)
	if err != nil {
		//Do something
		fmt.Println(err)
		os.Exit(1)
	}
}

func readFromFile(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		//Do something
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(content), "\n")

	for i, l := range lines {
		for j, v := range strings.Fields(l) {
			fmt.Println(i, j, v)
		}
	}
}
