package main

import "fmt"

func GetNthFib(n int) int {
	if n <= 1 {
		return 0
	} else if n == 2 {
		return 1
	}
	return GetNthFib(n-1) + GetNthFib(n-2)
}

func main() {
	input := 6
	fmt.Println("input: ", input)
	fmt.Println("output: ", GetNthFib(input))
}
