package main

import "fmt"

func OneEdit(stringOne string, stringTwo string) bool {
	lenDiff := len(stringOne) - len(stringTwo)
	if lenDiff > 1 || lenDiff < -1 {
		return false
	}
	if lenDiff == -1 {
		stringOne, stringTwo = stringTwo, stringOne
		lenDiff = 1
	}
	i := 0
	j := 0
	isEdit := false
	for i < len(stringOne) && j < len(stringTwo) {
		if stringOne[i] != stringTwo[j] {
			if isEdit {
				return false
			}
			isEdit = true
			if lenDiff == 0 {
				i++
				j++
			} else {
				i++
			}
		} else {
			i++
			j++
		}
	}
	return true
}

func main() {
	stringOne := "hello"
	stringTwo := "helo"
	fmt.Println("input: ", stringOne, stringTwo)
	fmt.Println("output: ", OneEdit(stringOne, stringTwo))
	stringOne = "bb"
	stringTwo = "a"
	fmt.Println("input: ", stringOne, stringTwo)
	fmt.Println("output: ", OneEdit(stringOne, stringTwo))
	stringOne = "a"
	stringTwo = "bb"
	fmt.Println("input: ", stringOne, stringTwo)
	fmt.Println("output: ", OneEdit(stringOne, stringTwo))
	stringOne = "hello"
	stringTwo = "elllos"
	fmt.Println("input: ", stringOne, stringTwo)
	fmt.Println("output: ", OneEdit(stringOne, stringTwo))
	stringOne = "helo"
	stringTwo = "helle"
	fmt.Println("input: ", stringOne, stringTwo)
	fmt.Println("output: ", OneEdit(stringOne, stringTwo))
}
