package main

import "fmt"

func GetNthFib(n int) int {
	// Write your code here.
	fibArr := make([]int, max(n, 2))
	fibArr[0] = 0
	fibArr[1] = 1
	for i := 2; i < n; i++ {
		fibArr[i] = fibArr[i-1] + fibArr[i-2]
	}
	return fibArr[n-1]
}

func main() {
	input := 6
	fmt.Println("input: ", input)
	fmt.Println("output: ", GetNthFib(input))
	input = 1
	fmt.Println("input: ", input)
	fmt.Println("output: ", GetNthFib(input))
}
