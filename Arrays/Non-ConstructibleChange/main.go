package main

import (
	"fmt"
	"sort"
)

func NonConstructibleChange(coins []int) int {
	// Write your code here.
	sort.Ints(coins)
	changeMade := 0
	for _, coin := range coins {
		if coin > changeMade+1 {
			return changeMade + 1
		}
		changeMade += coin
	}
	return changeMade + 1
}

func main() {
	fmt.Println([]int{5, 7, 1, 1, 2, 3, 22})
}
