package main

import "fmt"

type interval struct {
	left  int
	right int
}

func UnderscorifySubstring(str string, substring string) string {
	potentialMatch := map[int]interval{}
	byteStr := []byte(str)
	byteSubStr := []byte(substring)
	for i := 0; i < len(byteStr)-i; i++ {
		if str[i:i+len(substring)] == substring {
			potentialMatch[i] = interval{i, i + len(byteSubStr)}
		}
	}
	for k, v := range potentialMatch {
		left := v.left
		right := v.right
		for i := left; i <= right; i++ {
			if _, ok := potentialMatch[i]; ok {
				right = potentialMatch[i].right
				delete(potentialMatch, i)
			}
		}
		potentialMatch[k] = interval{left, right}
	}
	fmt.Println(potentialMatch)
	result := ""
	underscoreMap := map[int]bool{}
	for _, v := range potentialMatch {
		underscoreMap[v.left] = true
		underscoreMap[v.right] = true
	}
	fmt.Println(underscoreMap)
	for i, b := range byteStr {
		if _, matched := underscoreMap[i]; matched {
			result += "_"
		}
		result += string(b)
	}

	return result
}

func main() {
	str := "testthis is a testtest to see if testestest it works"
	substring := "test"
	fmt.Println(UnderscorifySubstring(str, substring))
}
