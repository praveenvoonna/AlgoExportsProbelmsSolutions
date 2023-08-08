package main

import "fmt"

func GenerateDocument(characters string, document string) bool {
	charMap := make(map[byte]int)
	for _, v := range characters {
		charMap[byte(v)] += 1
	}
	for _, v := range document {
		if charMap[byte(v)] == 0 {
			return false
		}
		charMap[byte(v)] -= 1
	}
	return true
}

func main() {
	fmt.Println(GenerateDocument("Bste!hetsi ogEAxpelrt x ", "AlgoExpert is the Best!"))
}
