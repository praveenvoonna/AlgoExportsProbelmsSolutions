package main

import "fmt"

func OneEdit(stringOne string, stringTwo string) bool {
	count := 0
	lenDiff := len(stringOne) - len(stringTwo)
	if lenDiff == 0 {
		for i := 0; i < len(stringOne); i++ {
			if stringOne[i] != stringTwo[i] {
				count++
			}
		}
	} else if lenDiff == 1 || lenDiff == -1 {
		if lenDiff == -1 {
			stringOne, stringTwo = stringTwo, stringOne
		}
		i := 0
		j := 0
		for i < len(stringOne) && j < len(stringTwo) {
			if stringOne[i] != stringTwo[j] {
				count++
				i++
				continue
			}
			i++
			j++
		}

	} else {
		return false
	}
	if count <= 1 {
		return true
	}
	return false
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
