package main

import "fmt"

func MinimumCharactersForWords(words []string) []string {
	freqMap := make(map[byte]int)
	for _, word := range words {
		charMap := make(map[byte]int)
		for _, c := range []byte(word) {
			charMap[c]++
		}
		for k, i := range charMap {
			if v, ok := freqMap[k]; ok {
				if v < i {
					freqMap[k] = i
				}
			} else {
				freqMap[k] = i
			}
		}
	}
	result := []string{}
	for k, v := range freqMap {
		for i := 0; i < v; i++ {
			result = append(result, string(k))
		}
	}
	return result
}

func main() {
	input := []string{"this", "that", "did", "deed", "them!", "a"}
	fmt.Println("input: ", input)
	fmt.Println("output: ", MinimumCharactersForWords(input))
}
