package main

import "fmt"

func InsertionSort(array []int) []int {
	for i := range array {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
	return array
}

func main() {
	fmt.Println(InsertionSort([]int{8, 5, 2, 9, 5, 6, 3}))
}
