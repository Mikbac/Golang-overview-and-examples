package main

import (
	"fmt"
)

// https://leetcode.com/problems/roman-to-integer/description/

func main() {
	fmt.Println(romanToInt("III"))     // 3
	fmt.Println(romanToInt("LVIII"))   // 58
	fmt.Println(romanToInt("MCMXCIV")) // 1994

	fmt.Println("---------------------------------")

}

// Runtime: 0 ms
// Memory: 4.49 MB
func romanToInt(s string) int {
	num := map[byte]int{
		73: 1,
		86: 5,
		88: 10,
		76: 50,
		67: 100,
		68: 500,
		77: 1000,
	}

	if len(s) == 1 {
		return num[s[0]]
	}

	sum := 0
	iterN := len(s) - 1

	for i := 0; i < iterN; i++ {
		curV := num[s[i]]
		if num[s[i+1]] > curV {
			sum -= curV
		} else {
			sum += curV
		}
	}
	sum += num[s[iterN]]

	return sum
}
