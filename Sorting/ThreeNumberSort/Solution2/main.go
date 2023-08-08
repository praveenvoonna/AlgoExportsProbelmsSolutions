package main

import "fmt"

func ThreeNumberSort(array []int, order []int) []int {
	firstOrder, secondOrder := order[0], order[1]
	firstIdx, secondIdx, thirdIdx := 0, 0, len(array)-1
	for secondIdx <= thirdIdx {
		value := array[secondIdx]
		if value == firstOrder {
			array[firstIdx], array[secondIdx] = array[secondIdx], array[firstIdx]
			firstIdx++
			secondIdx++
		} else if value == secondOrder {
			secondIdx++
		} else {
			array[secondIdx], array[thirdIdx] = array[thirdIdx], array[secondIdx]
			thirdIdx--
		}
	}
	return array
}

func main() {
	fmt.Println(ThreeNumberSort([]int{1, 0, 0, -1, -1, 0, 1, 1}, []int{0, 1, -1}))
}
