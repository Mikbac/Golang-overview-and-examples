package main

func runArraysDemo() {
	var sampleArray1 [3]int = [3]int{1, 2, 3}
	sampleArray2 := [3]int{1, 2, 3}

	for element := range sampleArray1 {
		println(element)
	}
	for element := range sampleArray2 {
		println(element)
	}
}
