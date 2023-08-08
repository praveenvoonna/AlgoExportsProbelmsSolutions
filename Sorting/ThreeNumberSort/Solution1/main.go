package main

import "fmt"

func ThreeNumberSort(array []int, order []int) []int {
	countOrder := []int{0, 0, 0}
	for i := 0; i < 3; i++ {
		for j := 0; j < len(array); j++ {
			if order[i] == array[j] {
				countOrder[i]++
			}
		}
	}
	j := 0
	for i := 0; i < 3; i++ {
		for countOrder[i] != 0 {
			array[j] = order[i]
			j++
			countOrder[i]--
		}
	}
	return array
}

func main() {
	fmt.Println(ThreeNumberSort([]int{1, 0, 0, -1, -1, 0, 1, 1}, []int{0, 1, -1}))
}
