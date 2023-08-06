package main

import "fmt"

func IsMonotonic(array []int) bool {
	if len(array) <= 2 {
		return true
	}
	tempVal := array[0]
	result := ""
	for _, val := range array {
		if tempVal >= val {
			tempVal = val
		} else {
			result = "not-decreasing"
			break
		}
	}
	if result == "" {
		return true
	}
	result = ""
	tempVal = array[0]
	for _, val := range array {
		if tempVal <= val {
			tempVal = val
		} else {
			result = "not-increasing"
			break
		}
	}
	if result == "" {
		return true
	}
	return false
}

func main() {
	fmt.Println(IsMonotonic([]int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}))
}
