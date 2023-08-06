package main

import (
	"fmt"
	"sort"
)

func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	targetPairs := [][]int{}
	for i := 0; i < len(array)-2; i++ {
		left, right := i+1, len(array)-1
		for left < right {
			currentSum := array[i] + array[left] + array[right]
			if currentSum == target {
				targetPairs = append(targetPairs, []int{array[i], array[left], array[right]})
				left += 1
				right -= 1
			} else if currentSum < target {
				left += 1
			} else {
				right -= 1
			}
		}
	}
	return targetPairs
}

func main() {
	fmt.Println(ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0))
}
