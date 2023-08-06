package main

import "fmt"

func TwoNumberSum(array []int, target int) []int {
	sumMap := make(map[int]int)
	for _, v := range array {
		desiredVal := target - v
		if v != desiredVal {
			if _, ok := sumMap[v]; !ok {
				sumMap[v] = 1
			}
			if _, ok := sumMap[desiredVal]; ok {
				return []int{v, desiredVal}
			}
		}
	}
	return []int{}
}

func main() {
	myArr := []int{1, 3, 4, 10, 5, 6}
	fmt.Println(TwoNumberSum(myArr, 7))
}
