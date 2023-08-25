package main

import "fmt"

func CommonCharacters(strings []string) []string {
	smallestString := getSmallestString(strings)
	potentialMatch := map[rune]bool{}
	for _, r := range smallestString {
		potentialMatch[r] = true
	}
	for _, str := range strings {
		uniqueChars := map[rune]bool{}
		for _, r := range str {
			uniqueChars[r] = true
		}
		for c := range potentialMatch {
			if _, ok := uniqueChars[c]; !ok {
				delete(potentialMatch, c)
			}
		}
	}
	result := []string{}
	for k := range potentialMatch {
		result = append(result, string(k))
	}
	return result
}

func getSmallestString(strings []string) string {
	smallestStr := strings[0]
	for _, str := range strings[1:] {
		if len(str) < len(smallestStr) {
			smallestStr = str
		}
	}
	return smallestStr
}

func main() {
	input := []string{"abc", "bcd", "cbad"}
	fmt.Println("input: ", input)
	fmt.Println("output: ", CommonCharacters(input))
}
