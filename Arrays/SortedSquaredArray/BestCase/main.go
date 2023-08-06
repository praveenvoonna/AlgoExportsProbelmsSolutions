package main

import "fmt"

func SortedSquaredArray(array []int) []int {
	// Write your code here.
	squaredArr := make([]int, len(array))
	leftP := 0
	rightP := len(array) - 1
	for i := rightP; i >= 0; i-- {
		rv := abs(array[rightP])
		lv := abs(array[leftP])
		if rv >= lv {
			squaredArr[i] = rv * rv
			rightP--
		} else {
			squaredArr[i] = lv * lv
			leftP++
		}
	}

	return squaredArr
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println([]int{1, 2, 3, 5, 6, 8, 9})
}
