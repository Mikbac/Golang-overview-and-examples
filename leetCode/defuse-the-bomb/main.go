package main

import (
	"fmt"
)

// https://leetcode.com/problems/defuse-the-bomb/description/

func main() {
	fmt.Println(decrypt([]int{5, 7, 1, 4}, 3))  // [12,10,16,13]
	fmt.Println(decrypt([]int{1, 2, 3, 4}, 0))  // [0,0,0,0]
	fmt.Println(decrypt([]int{2, 4, 9, 3}, -2)) // [12,5,6,13]
}

// Runtime: 0 ms
// Memory: 3.98 MB
func decrypt(code []int, k int) []int {
	cop := make([]int, len(code))
	codeLenIter := len(code)

	for i := 0; i < codeLenIter; i++ {
		sum := 0
		if k > 0 {
			for j := i + 1; j <= i+k; j++ {
				sum += code[j%codeLenIter]
			}
			cop[i] = sum
		} else {
			for j := i + k; j < i; j++ {
				if j < 0 {
					sum += code[codeLenIter-((j-j-j)%codeLenIter)]
				} else {
					sum += code[((j) % codeLenIter)]
				}
			}
		}
		cop[i] = sum
	}
	return cop
}
