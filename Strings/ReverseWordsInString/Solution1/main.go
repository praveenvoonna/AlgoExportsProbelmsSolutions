package main

import "fmt"

func ReverseWordsInString(str string) string {
	byteString := []byte(str)
	reverseByte := make([]byte, 0, len(byteString))
	for i := len(byteString) - 1; i >= 0; i-- {
		reverseByte = append(reverseByte, byteString[i])
	}
	wordByte := []byte{}
	finalByteStr := []byte{}
	for _, b := range reverseByte {
		if b != ' ' {
			wordByte = append(wordByte, b)
		} else {
			for i := len(wordByte) - 1; i >= 0; i-- {
				finalByteStr = append(finalByteStr, wordByte[i])
			}
			finalByteStr = append(finalByteStr, b)
			wordByte = []byte{}
		}
	}
	for i := len(wordByte) - 1; i >= 0; i-- {
		finalByteStr = append(finalByteStr, wordByte[i])
	}
	return string(finalByteStr)
}

func main() {
	fmt.Println(ReverseWordsInString("AlgoExpert is the best!"))
}
