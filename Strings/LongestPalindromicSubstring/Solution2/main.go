package main

import "fmt"

func LongestPalindromicSubstring(str string) string {
	if len(str) <= 1 {
		return str
	}
	longestStr := str[0:1]
	for i := 1; i < len(str); i++ {
		evenStr := MaxSubString(str, i-1, i)
		oddStr := MaxSubString(str, i-1, i+1)
		var maxStr string
		if len(oddStr) < len(evenStr) {
			maxStr = evenStr
		} else {
			maxStr = oddStr
		}
		if len(maxStr) > len(longestStr) {
			longestStr = maxStr
		}
	}
	return longestStr
}

func MaxSubString(str string, left, right int) string {
	maxStr := ""
	for left >= 0 && right < len(str) {
		if str[left] == str[right] {
			maxStr = str[left : right+1]
			left--
			right++
		} else {
			break
		}
	}
	return maxStr
}

func main() {
	input := "abaxyzzyxf"
	fmt.Println("input: ", input)
	fmt.Println("output: ", LongestPalindromicSubstring(input))
}
