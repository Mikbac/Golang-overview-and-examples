package main

import "fmt"

func runSliceDemo() {

	sampleSlice1 := []string{"aaa", "bbb"}
	sampleSlice2 := append(sampleSlice1, "ccc", "ddd")

	fmt.Println(sampleSlice1)

	for index, element := range sampleSlice2 {
		fmt.Println(index, element)
	}

}
