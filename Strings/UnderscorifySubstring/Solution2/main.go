package main

import "fmt"

type intervals []*interval

type interval struct {
	left  int
	right int
}

func UnderscorifySubstring(str string, substring string) string {
	locations := getLocations(str, substring)
	for _, location := range locations {
		fmt.Println(location)
	}
	fmt.Println("---------------------")
	collapsedLocation := colapseLocations(locations, substring)
	for _, location := range collapsedLocation {
		fmt.Println(location)
	}
	if len(collapsedLocation) < 1 {
		return str
	}
	resultByte := []byte{}
	locationIdx := 0
	for i, b := range []byte(str) {
		if collapsedLocation[locationIdx].left == i {
			resultByte = append(resultByte, '_')
		} else if collapsedLocation[locationIdx].right == i {
			resultByte = append(resultByte, '_')
			if locationIdx < len(collapsedLocation)-1 {
				locationIdx++
			}
		}
		resultByte = append(resultByte, b)
	}
	if collapsedLocation[locationIdx].right == len(str) {
		resultByte = append(resultByte, '_')
	}
	return string(resultByte)
}

func getLocations(str, substring string) intervals {
	intervals_temp := []*interval{}
	strLen := len(str)
	substringLen := len(substring)
	for i := 0; i < strLen-substringLen+1; i++ {
		if str[i:i+substringLen] == substring {
			intervalObj := interval{i, i + substringLen}
			intervals_temp = append(intervals_temp, &intervalObj)
		}
	}
	return intervals_temp
}

func colapseLocations(intervalsTemp intervals, substring string) intervals {
	if len(intervalsTemp) <= 1 {
		return intervalsTemp
	}
	instervalsResult := intervals{intervalsTemp[0]}
	previous := intervalsTemp[0]
	for i := 1; i < len(intervalsTemp); i++ {
		current := intervalsTemp[i]
		if previous.right >= current.left {
			previous.right = current.right
		} else {
			instervalsResult = append(instervalsResult, current)
			previous = current
		}
	}
	return instervalsResult
}

func main() {
	str := "testthis is a testtest to see if testestest it works"
	substring := "test"
	fmt.Println(UnderscorifySubstring(str, substring))
	str = "this is a test to see if it works and test"
	fmt.Println(UnderscorifySubstring(str, substring))
}
