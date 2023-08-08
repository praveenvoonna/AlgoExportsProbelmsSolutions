package main

import "fmt"

func HeapSort(array []int) []int {
	BuildMaxHeap(array)
	for i := len(array) - 1; i > 0; i-- {
		Swap(0, i, array)
		ShiftDown(0, i-1, array)
	}
	return array
}

func BuildMaxHeap(array []int) {
	lastParent := (len(array) - 2) / 2
	for i := lastParent; i >= 0; i-- {
		ShiftDown(i, len(array)-1, array)
	}
}

func ShiftDown(startIdx int, lastIdx int, array []int) {
	childOneIdx := startIdx*2 + 1
	childTwoIdx := -1
	for childOneIdx <= lastIdx {
		swapIdx := childOneIdx
		if childOneIdx+1 <= lastIdx {
			childTwoIdx = childOneIdx + 1
			if array[childOneIdx] < array[childTwoIdx] {
				swapIdx = childTwoIdx
			}
		}
		if array[startIdx] < array[swapIdx] {
			Swap(startIdx, swapIdx, array)
			startIdx = swapIdx
			childOneIdx = startIdx*2 + 1
		} else {
			return
		}
	}
}

func Swap(a, b int, array []int) {
	array[a], array[b] = array[b], array[a]
}

func main() {
	fmt.Println(HeapSort([]int{8, 5, 2, 9, 5, 6, 3}))
}
