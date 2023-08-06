package main

import "fmt"

func IsMonotonic(array []int) bool {
	isNonDecreasing := true
	isNonIncreasing := true
	for i := 1; i < len(array); i++ {
		if array[i-1] < array[i] {
			isNonDecreasing = false
		}
		if array[i-1] > array[i] {
			isNonIncreasing = false
		}
	}
	return isNonIncreasing || isNonDecreasing
}

func main() {
	fmt.Println(IsMonotonic([]int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}))
}
