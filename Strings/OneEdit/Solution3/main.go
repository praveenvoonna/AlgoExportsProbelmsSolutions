package main

import "fmt"

func OneEdit(stringOne string, stringTwo string) bool {
	lenOne := len(stringOne)
	lenTwo := len(stringTwo)
	lenDiff := lenOne - lenTwo
	if lenDiff > 1 || lenDiff < -1 {
		return false
	}

	for i := 0; i < min(lenOne, lenTwo); i++ {
		if stringOne[i] != stringTwo[i] {
			if lenOne < lenTwo {
				return stringOne[i:] == stringTwo[i+1:]
			} else if lenOne > lenTwo {
				return stringOne[i+1:] == stringTwo[i:]
			} else {
				return stringOne[i+1:] == stringTwo[i+1:]
			}
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
