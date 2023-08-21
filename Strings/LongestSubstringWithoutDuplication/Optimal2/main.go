package main

import "fmt"

type maxSubStr struct {
	left  int
	right int
}

func (m maxSubStr) length() int {
	return m.right - m.left
}

func LongestSubstringWithoutDuplication(str string) string {
	maxSubString := maxSubStr{0, 1}
	startIdx := 0
	lastSeenMap := map[rune]int{}
	for i, r := range str {
		if lastSeenIdx, found := lastSeenMap[r]; found {
			startIdx = max(startIdx, lastSeenIdx+1)
		}
		if maxSubString.length() < i+1-startIdx {
			maxSubString.left = startIdx
			maxSubString.right = i + 1
		}
		lastSeenMap[r] = i
	}
	return str[maxSubString.left:maxSubString.right]
}

func main() {
	input := "clementisacap"
	fmt.Println("input: ", input)
	fmt.Println("output: ", LongestSubstringWithoutDuplication(input))
}
