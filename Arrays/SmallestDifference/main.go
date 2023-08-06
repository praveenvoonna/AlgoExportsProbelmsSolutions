package main

import (
	"fmt"
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	smallest, currentVal := math.MaxInt32, math.MaxInt32
	idx1, idx2 := 0, 0
	smallestDiffer := []int{}
	for idx1 < len(array1) && idx2 < len(array2) {
		left, right := array1[idx1], array2[idx2]
		if left < right {
			currentVal = right - left
			idx1 += 1
		} else if right < left {
			currentVal = left - right
			idx2 += 1
		} else {
			return []int{array1[idx1], array2[idx2]}
		}
		if currentVal < smallest {
			smallest = currentVal
			smallestDiffer = []int{left, right}
		}
	}

	return smallestDiffer
}

func main() {
	fmt.Println(SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17}))
}
