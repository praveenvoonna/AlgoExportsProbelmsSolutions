package main

import (
	"fmt"
)

func Semordnilap(words []string) [][]string {
	semordnilap := map[string]bool{}
	result := [][]string{}
	for _, word := range words {
		semordnilap[word] = false
	}
	// fmt.Println(semordnilap)
	for word, _ := range semordnilap {
		reverseStr := reverseString(word)
		// fmt.Println(word, reverseStr)
		if _, ok := semordnilap[reverseStr]; ok && word != reverseStr {
			result = append(result, []string{word, reverseStr})
			delete(semordnilap, reverseStr)
		}
	}
	return result
}

func reverseString(word string) string {
	runeWord := []rune(word)
	for i, j := 0, len(runeWord)-1; i < j; i, j = i+1, j-1 {
		runeWord[i], runeWord[j] = runeWord[j], runeWord[i]
	}
	return string(runeWord)
}

func main() {
	input := []string{"desserts", "stressed", "hello"}
	fmt.Println("input: ", input)
	fmt.Println("output: ", Semordnilap(input))
}
