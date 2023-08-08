package main

import "fmt"

func CaesarCipherEncryptor(str string, key int) string {
	shift, offset := rune(key%26), rune(26)
	runes := []rune(str)
	for k, v := range runes {
		if v >= 'a' && v+shift <= 'z' {
			runes[k] = v + shift
		} else {
			runes[k] = v - offset + shift
		}
	}
	return string(runes)
}

func main() {
	fmt.Println(CaesarCipherEncryptor("xyz", 2))
}
