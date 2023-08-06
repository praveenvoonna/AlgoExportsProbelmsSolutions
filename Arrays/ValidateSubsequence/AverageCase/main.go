package main

import "fmt"

func IsValidSubsequence(array []int, sequence []int) bool {
	seqSize := len(sequence)
	arrSize := len(array)
	al := 0
	pl := 0
	sl := 0
	for sl < seqSize {
		for al < arrSize {
			if array[al] == sequence[sl] {
				al++
				sl++
			} else {
				al++
			}
		}
		if sl == seqSize {
			return true
		} else {
			sl = 0
			pl += 1
			al = pl
		}
		if pl > arrSize-seqSize {
			return false
		}
	}
	return true
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}
	fmt.Println(IsValidSubsequence(array, sequence))
}
