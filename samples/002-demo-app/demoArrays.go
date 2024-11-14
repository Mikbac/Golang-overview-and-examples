package main

import "fmt"

func runArraysDemo() {
	var sampleArray1 [3]int = [3]int{1, 2, 3}
	sampleArray2 := [3]int{1, 2, 3}

	for element := range sampleArray1 {
		fmt.Println(element)
	}
	for element := range sampleArray2 {
		fmt.Println(element)
	}
}
