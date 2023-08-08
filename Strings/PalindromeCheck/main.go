package main

import "fmt"

func IsPalindrome(str string) bool {
	left := 0
	right := len(str) - 1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("abcdcba"))
}
