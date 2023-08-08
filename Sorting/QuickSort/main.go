package main

import "fmt"

func QuickSort(array []int) []int {
	return QS(array, 0, len(array)-1)
}

func QS(array []int, left int, right int) []int {
	if left >= right {
		return array
	}
	pivot := Partition(array, left, right)
	QS(array, left, pivot-1)
	QS(array, pivot+1, right)
	return array
}
func Partition(array []int, left int, right int) int {
	pivot := array[left]
	i := left
	for j := left + 1; j <= right; j++ {
		if array[j] <= pivot {
			i += 1
			array[i], array[j] = array[j], array[i]
		}
	}
	array[left], array[i] = array[i], array[left]
	return i
}

func main() {
	fmt.Println(QuickSort([]int{8, 5, 2, 9, 5, 6, 3}))
}
