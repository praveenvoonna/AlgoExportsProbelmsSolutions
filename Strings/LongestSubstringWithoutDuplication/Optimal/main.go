package main

import "fmt"

func LongestSubstringWithoutDuplication(str string) string {
	locationMap := make(map[byte]int)
	maxSubStr := ""
	currSubStr := ""
	startIdx := 0
	// lastSeenIdx := 0
	for i, c := range []byte(str) {
		if v, ok := locationMap[c]; !ok {
			if len(currSubStr) == 0 {
				startIdx = i
			}
			locationMap[c] = i
			currSubStr += string(c)
			// fmt.Println("13 cur sub str", currSubStr, c, v)
		} else {
			locationMap[c] = i
			if len(currSubStr) > len(maxSubStr) {
				maxSubStr = currSubStr
			}
			// fmt.Println("19 cur sub str", currSubStr, c, v)
			startIdx = max(v+1, startIdx)
			currSubStr = str[startIdx : i+1]
			// fmt.Println("21 cur sub str", currSubStr, c, v)
		}
	}
	if len(currSubStr) > len(maxSubStr) {
		maxSubStr = currSubStr
	}
	return maxSubStr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := "clementisacap"
	fmt.Println("input: ", input)
	fmt.Println("output: ", LongestSubstringWithoutDuplication(input))
}
