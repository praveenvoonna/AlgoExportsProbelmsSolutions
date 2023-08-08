package main

import "fmt"

func GenerateDocument(characters string, document string) bool {
	charMap := make(map[byte]int)
	// docMap := make(map[byte]int)
	for _, v := range characters {
		if _, ok := charMap[byte(v)]; ok {
			charMap[byte(v)] += 1
		} else {
			charMap[byte(v)] = 1
		}
	}
	for _, v := range document {
		if c, ok := charMap[byte(v)]; ok && c > 0 {
			charMap[byte(v)] -= 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(GenerateDocument("Bste!hetsi ogEAxpelrt x ", "AlgoExpert is the Best!"))
}
