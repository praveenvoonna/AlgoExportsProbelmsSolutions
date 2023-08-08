package main

import "fmt"

func BinarySearch(array []int, target int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		middle := (left + right) / 2
		if array[middle] == target {
			return middle
		} else if array[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(BinarySearch([]int{1, 4, 6, 8, 10}, 10))
}
