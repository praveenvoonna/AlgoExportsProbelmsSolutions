package main

import (
	"fmt"
	"sort"
)

func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	resultArray := make([][]int, 0, 0)
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			for k := j + 1; k < len(array); k++ {
				if array[i]+array[j]+array[k] == target {
					resultItem := []int{array[i], array[j], array[k]}
					resultArray = append(resultArray, resultItem)
				}
			}
		}
	}
	return resultArray
}

func main() {
	fmt.Println(ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0))
}
