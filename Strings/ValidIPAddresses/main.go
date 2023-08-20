package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidIPAddresses(str string) []string {
	validIps := []string{}
	for i := 1; i < min(len(str), 4); i++ {
		ipAddress := []string{"", "", "", ""}
		ipAddress[0] = str[:i]
		// fmt.Println(ipAddress)
		if !isValid(ipAddress[0]) {
			// fmt.Println("16: ", ipAddress, ipAddress[0], isValid(ipAddress[0]))
			continue
		}
		for j := i + 1; j < i+min(len(str)-i, 4); j++ {
			ipAddress[1] = str[i:j]
			// fmt.Println(ipAddress)
			if !isValid(ipAddress[1]) {
				// fmt.Println("17", ipAddress, ipAddress[1], isValid(ipAddress[1]))
				continue
			}
			for k := j + 1; k < j+min(len(str)-j, 4); k++ {
				ipAddress[2] = str[j:k]
				ipAddress[3] = str[k:]
				// fmt.Println(ipAddress)
				if isValid(ipAddress[2]) && isValid(ipAddress[3]) {
					// fmt.Println("31", ipAddress, ipAddress[2], isValid(ipAddress[2]))
					// fmt.Println(ipAddress, ipAddress[3], isValid(ipAddress[3]))
					validIps = append(validIps, strings.Join(ipAddress, "."))
					// fmt.Println(ipAddress)
				}
			}
		}
	}
	return validIps
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isValid(str string) bool {
	i, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	if i > 255 {
		return false
	}
	return len(str) == len(strconv.Itoa(i))
}

func main() {
	// fmt.Println(isValid("1"))
	fmt.Println(ValidIPAddresses("1921680"))
}
