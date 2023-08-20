package main

import (
	"fmt"
	"strings"
)

func ReverseWordsInString(str string) string {
	strList := []string{}
	start := 0
	for i, c := range str {
		if c == ' ' {
			strList = append(strList, str[start:i])
			start = i
		} else if str[start] == ' ' {
			strList = append(strList, " ")
			start = i
		}
	}
	strList = append(strList, str[start:])
	reverseList(strList)
	return strings.Join(strList, "")
}

func reverseList(list []string) {
	start := 0
	end := len(list) - 1
	for start < end {
		list[start], list[end] = list[end], list[start]
		start++
		end--
	}

}

func main() {
	fmt.Println(ReverseWordsInString("AlgoExpert is the best!"))
}
