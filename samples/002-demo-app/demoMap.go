package main

import "fmt"

func runMapDemo() {
	values1 := map[string]int{
		"aaa": 1,
		"bbb": 2,
	}
	fmt.Println(values1) // map[aaa:1 bbb:2]

	var values2 map[string]int
	fmt.Println(values2) // map[]

	values3 := make(map[string]int)
	values3["ccc"] = 3
	values3["ddd"] = 4
	fmt.Println(values3)        // map[ccc:3 ddd:4]
	fmt.Println(values3["ccc"]) // 3
	delete(values3, "ddd")
	fmt.Println(values3) // map[ccc:3]

	values4 := map[string]int{
		"aaa": 1,
		"bbb": 2,
		"ccc": 3,
	}
	for key, value := range values4 {
		fmt.Println(key, value)
	}

}
