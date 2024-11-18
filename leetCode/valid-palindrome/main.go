package main

import (
	"fmt"
	"strings"
	"unicode"
)

// https://leetcode.com/problems/valid-palindrome/description/

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome(" "))                              // true
}

// Runtime: 0 ms
// Memory: 4.5 MB
func isPalindrome(s string) bool {
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}

		return unicode.ToLower(r)
	}
	s = strings.Map(f, s)
	// it makes the same, but it is less optimal:
	//s = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(strings.ToLower(s), "")

	sLen := len(s)
	iterN := sLen / 2
	for i := 0; i < iterN; i++ {
		if s[i] != s[sLen-1-i] {
			return false
		}
	}
	return true
}
