package main

func runSliceDemo() {

	sampleSlice1 := []string{"aaa", "bbb"}
	sampleSlice2 := append(sampleSlice1, "ccc", "ddd")

	println(sampleSlice1)

	for index, element := range sampleSlice2 {
		println(index, element)
	}

}
