package main

import (
	"fmt"
	"strconv"
)

func RunLengthEncoding(str string) string {
	lastCount := 1
	outputStr := make([]byte, 0, len(str))
	for i := 1; i < len(str); i++ {
		currentChar := str[i]
		prevChar := str[i-1]
		if currentChar != prevChar || lastCount == 9 {
			outputStr = append(outputStr, strconv.Itoa(lastCount)[0], prevChar)
			lastCount = 0
		}
		lastCount += 1
	}
	if lastCount > 0 && len(str) > 0 {
		outputStr = append(outputStr, strconv.Itoa(lastCount)[0], str[len(str)-1])
	}
	return string(outputStr)
}

func main() {
	fmt.Println(RunLengthEncoding("AAAAAAAAAAAAABBCCCCDD"))
}
