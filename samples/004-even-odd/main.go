package main

import "fmt"

func main() {
	numbers := []int{}
	for i := 0; i < 11; i++ {
		numbers = append(numbers, i)
	}

	for _, element := range numbers {
		if element%2 == 0 {
			fmt.Printf("Number %v is even\n", element)
		} else {
			fmt.Printf("Number %v is odd\n", element)
		}

	}
}
