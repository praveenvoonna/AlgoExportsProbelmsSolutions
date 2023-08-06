package main

import "fmt"

func TwoNumberSum(array []int, target int) []int {
	arrSize := len(array)
	for i := 0; i < arrSize; i++ {
		for j := i + 1; j < arrSize; j++ {
			if array[i]+array[j] == target {
				return []int{array[i], array[j]}
			}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(TwoNumberSum([]int{4, 6, 9, 1}, 7))
}
