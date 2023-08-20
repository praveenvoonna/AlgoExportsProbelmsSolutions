package main

import (
	"fmt"
	"sort"
)

func GroupAnagrams(words []string) [][]string {
	wordMap := make(map[string][]string)
	for _, word := range words {
		sortedWord := sortWord(word)
		wordMap[sortedWord] = append(wordMap[sortedWord], word)
	}
	result := make([][]string, 0, len(wordMap))
	for _, v := range wordMap {
		result = append(result, v)
	}
	return result
}

func sortWord(word string) string {
	byteWord := []byte(word)
	sort.Slice(byteWord, func(i, j int) bool {
		return byteWord[i] < byteWord[j]
	})
	return string(byteWord)
}

func main() {
	fmt.Println(GroupAnagrams([]string{"yo", "act", "flop", "tac", "foo", "cat", "oy", "olfp"}))
}
