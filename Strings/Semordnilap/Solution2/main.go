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
		// fmt.Println(word, reverseStr)s
		if _, ok := semordnilap[reverseStr]; ok && word != reverseStr {
			result = append(result, []string{word, reverseStr})
			delete(semordnilap, reverseStr)
		}
	}
	return result
}

func reverseString(word string) string {
	byteWord := make([]byte, 0, len(word))
	for i := len(word) - 1; i >= 0; i-- {
		byteWord = append(byteWord, word[i])
	}
	return string(byteWord)
}

func main() {
	input := []string{"desserts", "stressed", "hello"}
	fmt.Println("input: ", input)
	fmt.Println("output: ", Semordnilap(input))
	input = []string{"dog", "god"}
	fmt.Println("input: ", input)
	fmt.Println("output: ", Semordnilap(input))
}
