package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/palindrome-number/description/

func main() {
	fmt.Println(isPalindrome1(1221))  // true
	fmt.Println(isPalindrome1(12321)) // true
	fmt.Println(isPalindrome1(1))     // true
	fmt.Println(isPalindrome1(-121))  // false
	fmt.Println(isPalindrome1(121))   // true
	fmt.Println(isPalindrome1(10))    // false

	fmt.Println("---------------------------------")

	fmt.Println(isPalindrome2(1221))  // true
	fmt.Println(isPalindrome2(12321)) // true
	fmt.Println(isPalindrome2(1))     // true
	fmt.Println(isPalindrome2(-121))  // false
	fmt.Println(isPalindrome2(121))   // true
	fmt.Println(isPalindrome2(10))    // false
}

// Runtime: 0 ms
// Memory: 6 MB
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	sLen := len(s)
	iterN := sLen / 2
	for i := 0; i < iterN; i++ {
		if s[i] != s[sLen-1-i] {
			return false
		}
	}
	return true
}

// Runtime: 10 ms
// Memory: 5.9 MB
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	for x > 0 {
		// 123:
		// 0 + 3 = 3
		// 30 + 2 = 32
		// 320 + 1 = 321
		reversed = reversed*10 + x%10
		x = x / 10
	}

	return original == reversed
}
