package main

import (
	"fmt"
	"math"
)

func FindThreeLargestNumbers(array []int) []int {
	largestInts := []int{math.MinInt32, math.MinInt32, math.MinInt32}
	for i := 0; i < len(array); i++ {
		InsertIntoLargestArray(largestInts, array[i])
	}
	return largestInts
}

func InsertIntoLargestArray(largestInts []int, num int) {
	if num > largestInts[2] {
		ShiftAndUpdate(largestInts, num, 2)
	} else if num > largestInts[1] {
		ShiftAndUpdate(largestInts, num, 1)
	} else if num > largestInts[0] {
		ShiftAndUpdate(largestInts, num, 0)
	}
}

func ShiftAndUpdate(largestInts []int, num int, position int) {
	for i := 0; i < position+1; i++ {
		if position == i {
			largestInts[i] = num
		} else {
			largestInts[i] = largestInts[i+1]
		}
	}
}

func main() {
	fmt.Println(FindThreeLargestNumbers([]int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}))
}
