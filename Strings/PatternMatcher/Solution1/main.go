package main

import "fmt"

func PatternMatcher(pattern, str string) []string {
	patternCount := map[rune]int{}
	for _, r := range pattern {
		patternCount[r]++
	}

	// fmt.Println("pattern count: ", patternCount)
	for i := 1; i < len(str); i++ {
		if len(patternCount) == 2 {
			xSize := i
			ySize := (len(str) - (xSize * patternCount['x'])) / patternCount['y']
			if ySize >= 1 {
				patternSize := map[rune]int{'x': xSize, 'y': ySize}
				// fmt.Println(patternSize)
				if v, found := isValidPattern(pattern, str, patternSize); found {
					return v
				}
			}
		} else if len(patternCount) == 1 {
			isValid := true
			// firstPattern := ""
			patternNew := map[rune]string{}
			for k, v := range patternCount {
				xSize := i
				// fmt.Println("xSize: ", xSize)
				firstPattern := ""
				if xSize*v == len(str) {
					firstPattern = str[i : i+v]
					// fmt.Println("first pattern: ", firstPattern)
					for j := i + v; j < len(str)-v; j += v {
						curPattern := str[j : j+v]
						// fmt.Println("cur pattern: ", curPattern)
						if firstPattern != curPattern {
							isValid = false
							break
						}
					}
				} else {
					isValid = false
				}
				if isValid {
					patternNew[k] = firstPattern
					return []string{patternNew['x'], patternNew['y']}
				}
			}

		} else {
			return []string{}
		}
	}

	return []string{}
}

func isValidPattern(pattern, str string, patternSize map[rune]int) ([]string, bool) {
	// fmt.Println(pattern, str, patternSize)
	i := 0
	patternNew := map[rune]string{}
	for _, p := range pattern {
		if i+patternSize[p] > len(str) {
			return nil, false
		}
		// fmt.Println("i: ", i, "p: ", string(p))
		if v, found := patternNew[p]; !found {
			patternNew[p] = str[i : i+patternSize[p]]
			// fmt.Println("x: ", patternNew)
		} else {
			if v != str[i:i+patternSize[p]] {
				return nil, false
			}
		}
		i += patternSize[p]
		// fmt.Println("i: ", i)
		// fmt.Println(patternNew)
	}
	return []string{patternNew['x'], patternNew['y']}, true
}

func main() {
	pattern := "xxyxxy"
	str := "gogopowerrangergogopowerranger"
	fmt.Println(PatternMatcher(pattern, str))
	pattern = "xyx"
	str = "thisshouldobviouslybewrong"
	fmt.Println(PatternMatcher(pattern, str))
	pattern = "yxyx"
	str = "abab"
	fmt.Println(PatternMatcher(pattern, str))
	pattern = "xxxx"
	str = "testtesttesttest"
	fmt.Println(PatternMatcher(pattern, str))
}
