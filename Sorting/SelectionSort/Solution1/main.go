package main

import "fmt"

func SelectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		num := array[i]
		index := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < num {
				num = array[j]
				index = j
			}
		}
		array[i], array[index] = array[index], array[i]

	}
	return array
}

func main() {
	fmt.Println(SelectionSort([]int{8, 5, 2, 9, 5, 6, 3}))
}
