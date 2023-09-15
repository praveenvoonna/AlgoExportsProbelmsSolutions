package main

import (
	"fmt"
	"math"
)

func LongestBalancedSubstring(str string) int {
	boundaries := []int{0, math.MaxInt32}
	rightIdx := 0
	leftIdx := 0
	leftPar := 0

	for rightIdx < len(str) {
		if str[rightIdx] == '(' {
			leftPar++
		} else if str[rightIdx] == ')' {
			if leftPar >= 0 {
				leftPar--
			} else {
				boundaries = maxBoundaries(boundaries, leftIdx, rightIdx)
				leftIdx++
				rightIdx = leftIdx
				leftPar = 0
				fmt.Println(boundaries, leftIdx, rightIdx)
			}
		}
		rightIdx++
	}
	fmt.Println(leftIdx, rightIdx)
	boundaries = maxBoundaries(boundaries, leftIdx, rightIdx)
	return getMaxStrCount(str, boundaries)
}

func maxBoundaries(boundaries []int, left, right int) []int {
	if boundaries[1] == math.MaxInt32 {
		return []int{left, right}
	}
	if boundaries[1]-boundaries[0] < right-left {
		return []int{left, right}
	}
	return boundaries
}

func getMaxStrCount(str string, boundaries []int) int {
	if boundaries[1] == math.MaxInt32 {
		return -1
	}
	return boundaries[1] - boundaries[0]
}

func main() {
	input := "(()))("
	fmt.Println("input: ", input)
	fmt.Println("output: ", LongestBalancedSubstring(input))
	input = "())()(()())"
	fmt.Println("input: ", input)
	fmt.Println("output: ", LongestBalancedSubstring(input))
}
