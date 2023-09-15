package main

import (
	"fmt"
	"math"
)

func LongestBalancedSubstring(str string) int {
	boundaries := []int{0, math.MaxInt32}
	for i := 0; i < len(str); i++ {
		for j := i + 2; j <= len(str); j++ {
			if isBalanced(str, i, j) {
				// fmt.Println(str[i:j], i, j)
				boundaries = getMaxBoundaries(boundaries, i, j)
			}
		}
	}
	return getMaxCount(boundaries)
}

func isBalanced(str string, left, right int) bool {
	leftPar := 0
	// fmt.Println(str[left:right])
	for i := left; i < right; i++ {
		if str[i] == '(' {
			leftPar++
		} else if leftPar > 0 {
			leftPar--
		} else {
			return false
		}
	}
	return leftPar == 0
}

func getMaxBoundaries(boundaries []int, i, j int) []int {
	if boundaries[1] == math.MaxInt32 {
		return []int{i, j}
	}
	if boundaries[1]-boundaries[0] < j-i {
		return []int{i, j}
	}
	return boundaries
}

func getMaxCount(boundaries []int) int {
	if boundaries[1] == math.MaxInt32 {
		return 0
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

	input = "())()(()())"
	fmt.Println("input: ", input)
	fmt.Println("output: ", LongestBalancedSubstring(input))

	input = "(("
	fmt.Println("input: ", input)
	fmt.Println("output: ", LongestBalancedSubstring(input))

}
