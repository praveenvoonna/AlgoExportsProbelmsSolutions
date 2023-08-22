package main

import (
	"fmt"
	"strings"
)

func PatternMatcher(pattern, str string) []string {
	newPettrn, isSwapped := changePattern(pattern)
	// fmt.Println(newPettrn, isSwapped)
	var counts struct {
		x int
		y int
	}
	firstYPos := getCounts(newPettrn, &counts)
	if counts.y != 0 {
		for lenX := 1; lenX < len(str); lenX++ {
			lenY := (len(str) - lenX*counts.x) / counts.y
			if lenY < 1 || lenX*counts.x+lenY*counts.y != len(str) {
				continue
			}
			// fmt.Println(lenX, lenY, len(str), newPettrn)
			xString := str[:lenX]
			yString := str[firstYPos*lenX : firstYPos*lenX+lenY]
			// fmt.Println(xString, yString)
			patternStr := ""
			for _, r := range newPettrn {
				if r == 'x' {
					patternStr += xString
				} else {
					patternStr += yString
				}
			}
			if patternStr == str {
				if !isSwapped {
					return []string{xString, yString}
				} else {
					return []string{yString, xString}
				}
			}
		}
	} else {
		xLen := len(str) / counts.x
		xString := str[:xLen]
		patternStr := strings.Repeat(xString, xLen)
		if patternStr == str {
			if !isSwapped {
				return []string{xString, ""}
			} else {
				return []string{"", xString}
			}
		}
	}
	return []string{}
}

func changePattern(pattern string) (string, bool) {
	isSwapped := false
	if len(pattern) < 1 {
		return pattern, isSwapped
	}
	if pattern[0] == 'y' {
		isSwapped = true
	}
	newPattern := ""
	if isSwapped {
		for _, r := range pattern {
			if r == 'x' {
				newPattern += "y"
			} else {
				newPattern += "x"
			}
		}
		pattern = newPattern
	}
	return pattern, isSwapped
}

func getCounts(pattern string, counts *struct {
	x int
	y int
}) int {
	yPos := -1
	for i, p := range pattern {
		if p == 'x' {
			counts.x++
		} else {
			counts.y++
			if yPos == -1 {
				yPos = i
			}
		}
	}
	return yPos
}

func main() {
	pattern := "xxyxxy"
	str := "gogopowerrangergogopowerranger"
	fmt.Println("Pattern: ", pattern, "String: ", str)
	fmt.Println("Output: ", PatternMatcher(pattern, str))
	pattern = "xyx"
	str = "thisshouldobviouslybewrong"
	fmt.Println("Pattern: ", pattern, "String: ", str)
	fmt.Println("Output: ", PatternMatcher(pattern, str))
	pattern = "yxyx"
	str = "abab"
	fmt.Println("Pattern: ", pattern, "String: ", str)
	fmt.Println("Output: ", PatternMatcher(pattern, str))
	pattern = "xxxx"
	str = "testtesttesttest"
	fmt.Println("Pattern: ", pattern, "String: ", str)
	fmt.Println("Output: ", PatternMatcher(pattern, str))
}
