package main

import "fmt"

func LongestPalindromicSubstring(str string) string {
	palindrome := ""
	for i := 0; i < len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			checkStr := str[i:j]
			if len(checkStr) > len(palindrome) && isPalindrome(checkStr) {
				palindrome = checkStr
			}
		}
	}
	return palindrome
}

func isPalindrome(str string) bool {
	i, j := 0, len(str)-1
	for i < j {
		if str[i] == str[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func main() {
	input := "abaxyzzyxf"
	fmt.Println("input: ", input)
	fmt.Println("output: ", LongestPalindromicSubstring(input))
}
