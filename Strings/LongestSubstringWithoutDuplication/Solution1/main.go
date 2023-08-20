package main

import "fmt"

func LongestSubstringWithoutDuplication(str string) string {
	if len(str) <= 1 {
		return str
	}
	longestSubStr := ""
	for i := 0; i < len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			curStr := str[i:j]
			if len(longestSubStr) < len(curStr) && isDuplicate(curStr) {
				longestSubStr = curStr
			}
		}
	}
	return longestSubStr
}

func isDuplicate(str string) bool {
	freqMap := make(map[rune]bool)
	for _, r := range str {
		if freqMap[r] {
			return false
		} else {
			freqMap[r] = true
		}
	}
	return true
}

func main() {
	input := "clementisacap"
	fmt.Println("input: ", input)
	fmt.Println("output: ", LongestSubstringWithoutDuplication(input))
}
