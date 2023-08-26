package main

import (
	"fmt"
	"math"
)

func SmallestSubstringContaining(bigString, smallString string) string {
	charCounts := getCharCounts(smallString)
	subStringBounds := getSubstringBounds(bigString, charCounts)
	return getSubString(bigString, subStringBounds)
}

func getCharCounts(str string) map[byte]int {
	charCounts := map[byte]int{}
	for _, r := range str {
		increaseMapCounts(charCounts, byte(r))
	}
	return charCounts
}

func getSubstringBounds(str string, charCounts map[byte]int) []int {
	subStringBounds := []int{0, math.MaxInt32}
	leftIdx := 0
	rightIdx := 0
	uniqueCharCount := map[byte]int{}
	uniqueCharCountDone := 0
	charCountsLen := len(charCounts)
	for rightIdx < len(str) {
		rightChar := str[rightIdx]
		if _, ok := charCounts[rightChar]; !ok {
			rightIdx++
			continue
		}
		increaseMapCounts(uniqueCharCount, rightChar)
		if uniqueCharCount[rightChar] == charCounts[rightChar] {
			uniqueCharCountDone++
		}

		for uniqueCharCountDone == charCountsLen && leftIdx <= rightIdx {
			leftChar := str[leftIdx]
			subStringBounds = getLeastCharStrBounds(subStringBounds, leftIdx, rightIdx)
			if _, ok := charCounts[leftChar]; !ok {
				leftIdx++
				continue
			}
			if uniqueCharCount[leftChar] == charCounts[leftChar] {
				uniqueCharCountDone--
			}
			decreaseMapCounts(uniqueCharCount, leftChar)
			leftIdx++
		}

		rightIdx++
	}
	return subStringBounds
}

func getLeastCharStrBounds(subStringBounds []int, left, right int) []int {
	if subStringBounds[1]-subStringBounds[0] > right-left {
		return []int{left, right}
	}
	return subStringBounds
}

func getSubString(str string, bounds []int) string {
	if bounds[1] == math.MaxInt32 {
		return ""
	}
	return str[bounds[0] : bounds[1]+1]
}

func increaseMapCounts(charCounts map[byte]int, b byte) {
	charCounts[b]++
}

func decreaseMapCounts(charCounts map[byte]int, b byte) {
	charCounts[b]--
}

func main() {
	bigString := "abcd$ef$axb$c$"
	smallString := "$$abf"
	fmt.Println("input: ", "bigString: ", bigString, " smallString: ", smallString)
	fmt.Println("output: ", SmallestSubstringContaining(bigString, smallString))
	bigString = "14562435612z!8828!193236!336!5$41!23!5!6789901#z2!"
	smallString = "#!2z"
	fmt.Println("input: ", "bigString: ", bigString, " smallString: ", smallString)
	fmt.Println("output: ", SmallestSubstringContaining(bigString, smallString))
}
