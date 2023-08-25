package main

import "fmt"

func CommonCharacters(strings []string) []string {
	charCounts := map[rune]int{}
	for _, str := range strings {
		uniqueChars := map[rune]bool{}
		for _, b := range str {
			uniqueChars[b] = true
		}
		for k := range uniqueChars {
			charCounts[k]++
		}

	}
	result := []string{}
	for k, v := range charCounts {
		if v == len(strings) {
			result = append(result, string(k))
		}
	}
	return result
}

func main() {
	input := []string{"abc", "bcd", "cbad"}
	fmt.Println("input: ", input)
	fmt.Println("output: ", CommonCharacters(input))
}
