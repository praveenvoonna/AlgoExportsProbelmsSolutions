package main

import (
	"fmt"
	"sort"
)

func SortedSquaredArray(array []int) []int {
	// Write your code here.
	for k, v := range array {
		array[k] = v * v
	}
	sort.Ints(array)
	return array
}

func main() {
	fmt.Println([]int{1, 2, 3, 5, 6, 8, 9})
}
