package main

import "fmt"

func MoveElementToEnd(array []int, toMove int) []int {
	left, right := 0, len(array)-1
	for left < right {
		for array[right] == toMove && left < right {
			right -= 1
		}
		if array[left] == toMove {
			array[left] = array[right]
			array[right] = toMove
			left += 1

		} else {
			left += 1
		}
	}
	return array
}

func main() {
	fmt.Println(MoveElementToEnd([]int{2, 1, 2, 2, 2, 3, 4, 2}, 2))
}
