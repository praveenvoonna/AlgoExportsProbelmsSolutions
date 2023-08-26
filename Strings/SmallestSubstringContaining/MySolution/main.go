package main

import "fmt"

func SmallestSubstringContaining(bigString, smallString string) string {
	smallStrChars := map[byte][]int{}
	for _, b := range []byte(smallString) {
		if _, ok := smallStrChars[b]; !ok {
			smallStrChars[b] = []int{1, 0}
		} else {
			smallStrChars[b][0]++
		}
	}
	fmt.Println("smallStrChars: ", smallStrChars)
	subStr := ""
	i := 0
	smallLen := len(smallString)
	bigLen := len(bigString)
	fmt.Println("subStr: ", subStr, smallLen, bigLen)
	for i < bigLen {
		r := bigString[i]
		if _, ok := smallStrChars[r]; ok && i+smallLen <= bigLen {
			for k := range smallStrChars {
				smallStrChars[k][1] = 0
			}
			fmt.Println(i, r, smallStrChars)
			j := i
			strSize := 0
			for j < bigLen {
				b2 := bigString[j]
				if v, found := smallStrChars[b2]; found && v[0] > v[1] {
					smallStrChars[b2][1]++
					strSize++
				}
				if strSize == smallLen {
					newStr := bigString[i : j+1]
					if len(subStr) == 0 || len(subStr) > len(newStr) {
						subStr = newStr
					}
				}
				j++
			}

		}
		i++
	}

	return subStr
}

func main() {
	bigString := "abcd$ef$axb$c$"
	smallString := "$$abf"
	fmt.Println("input: ", "bigString: ", bigString, " smallString: ", smallString)
	fmt.Println("output: ", SmallestSubstringContaining(bigString, smallString))
	bigString = "14562435612z!8828!193236!336!5$41!23!5!6789901#z2!"
	smallString = "#!2z"
	fmt.Println("input: ", "bigString: ", bigString, " smallString: ", smallString)
	fmt.Println("output: ", SmallestSubstringContaining(bigString, smallString))
}
