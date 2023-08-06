package main

import "fmt"

func IsValidSubsequence(array []int, sequence []int) bool {
	arrP := 0
	seqP := 0
	for arrP < len(array) && seqP < len(sequence) {
		if array[arrP] == sequence[seqP] {
			seqP++
		}
		arrP++
	}

	return seqP == len(sequence)
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}
	fmt.Println(IsValidSubsequence(array, sequence))
}
