package main

import "fmt"

func ThreeNumberSort(array []int, order []int) []int {
	firstValue, thirdValue := order[0], order[2]
	firstIdx := 0
	for i := 0; i < len(array); i++ {
		if array[i] == firstValue {
			array[i], array[firstIdx] = array[firstIdx], array[i]
			firstIdx++
		}
	}
	thirdIdx := len(array) - 1
	for i := thirdIdx; i >= firstIdx; i-- {
		if array[i] == thirdValue {
			array[i], array[thirdIdx] = array[thirdIdx], array[i]
			thirdIdx--
		}
	}
	return array
}

func main() {
	fmt.Println(ThreeNumberSort([]int{1, 0, 0, -1, -1, 0, 1, 1}, []int{0, 1, -1}))
}
