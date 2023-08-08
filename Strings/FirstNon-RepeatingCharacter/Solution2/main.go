package main

import "fmt"

func FirstNonRepeatingCharacter(str string) int {
	charFreq := make(map[rune]int)
	for _, v := range str {
		charFreq[v] += 1
	}
	for i, v := range str {
		if charFreq[v] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(FirstNonRepeatingCharacter("abcdcaf"))
}
