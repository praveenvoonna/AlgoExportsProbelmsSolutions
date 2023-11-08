package main

import (
	"fmt"
	"math"
)

func RadixSort(array []int) []int {
	maxNum := math.MinInt32
	for _, num := range array {
		if num > maxNum {
			maxNum = num
		}
	}
	i := 1
	for maxNum > 0 {
		maxNum /= 10
		array = RS(array, i)
		i *= 10
	}
	return array
}

func RS(array []int, digit int) []int {
	sorted := make([]int, len(array))
	counts := make([]int, 10)
	for _, num := range array {
		counts[(num/digit)%10]++
	}
	for i := 1; i < 10; i++ {
		counts[i] += counts[i-1]
	}
	for i := len(array) - 1; i >= 0; i-- {
		tempNum := array[i] / digit
		place := tempNum % 10
		counts[place]--
		sorted[counts[place]] = array[i]
	}
	return sorted
}

func main() {
	fmt.Println(RadixSort([]int{8762, 654, 3008, 345, 87, 65, 234, 12, 2}))
}
