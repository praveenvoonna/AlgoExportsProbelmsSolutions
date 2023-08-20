package main

import (
	"fmt"
	"sort"
)

func GroupAnagrams(words []string) [][]string {
	if len(words) == 0 {
		return [][]string{}
	}
	sortedWords := make([]string, 0, len(words))
	indices := make([]int, 0, len(words))
	for i, word := range words {
		sortedWords = append(sortedWords, sortWord(word))
		indices = append(indices, i)
	}
	sort.Slice(indices, func(i, j int) bool {
		return sortedWords[indices[i]] < sortedWords[indices[j]]
	})

	result := [][]string{}
	currentAnagramGroup := []string{}
	currentAnagram := sortedWords[indices[0]]
	for _, index := range indices {
		word := words[index]
		sortedWord := sortedWords[index]
		if len(currentAnagramGroup) == 0 {
			currentAnagramGroup = append(currentAnagramGroup, word)
			currentAnagram = sortedWord
			continue
		}

		if sortedWord == currentAnagram {
			currentAnagramGroup = append(currentAnagramGroup, word)
		} else {
			result = append(result, currentAnagramGroup)
			currentAnagramGroup = []string{word}
			currentAnagram = sortedWord
		}
	}
	result = append(result, currentAnagramGroup)
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
