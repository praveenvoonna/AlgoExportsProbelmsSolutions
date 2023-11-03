package main

import "fmt"

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	// Write your code here.
	return helper(array, 1)
}

func helper(array []interface{}, product int) int {
	sum := 0
	for _, elem := range array {
		if cast, ok := elem.(SpecialArray); ok {
			sum += helper(cast, product+1)
		} else if cast, ok := elem.(int); ok {
			sum += cast
		}

	}
	return product * sum
}

func main() {
	input := SpecialArray{
		5, 2,
		SpecialArray{7, -1},
		3,
		SpecialArray{
			6,
			SpecialArray{
				-13, 8,
			},
			4,
		},
	}
	fmt.Println("input: ", input)
	fmt.Println("output: ", ProductSum(input))
}
