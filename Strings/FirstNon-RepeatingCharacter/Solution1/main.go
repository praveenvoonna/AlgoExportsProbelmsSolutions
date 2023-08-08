package main

import "fmt"

func FirstNonRepeatingCharacter(str string) int {
	if len(str) == 1 {
		return 0
	}
	type position struct {
		isMultiple bool
		place      int
	}
	runeMap := make(map[rune]position)
	for i, r := range str {
		pos := position{}
		if _, ok := runeMap[r]; ok {
			pos.isMultiple = false
		} else {
			pos.isMultiple = true
			pos.place = i
		}
		runeMap[r] = pos
	}
	pos := len(str) - 1
	isUpdated := false
	for _, v := range runeMap {
		if v.isMultiple {
			if v.place <= pos {
				pos = v.place
				isUpdated = true
			}
		}
	}
	if !isUpdated {
		return -1
	}
	return pos
}

func main() {
	fmt.Println(FirstNonRepeatingCharacter("abcdcaf"))
}
