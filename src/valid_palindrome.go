package main

import (
	"fmt"
	strings "strings"
)

func main() {
	s := ""
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = removeNonAlphanumeric(s)
	backwardS := getBackwardS(s)
	return strings.EqualFold(s, backwardS)
}

func removeNonAlphanumeric(s string) string {
	var alphanumericS strings.Builder

	for _, symbol := range s {
		if (symbol >= 48 && symbol <= 57) || (symbol >= 97 && symbol <= 122) {
			alphanumericS.WriteRune(symbol)
		}
	}

	return alphanumericS.String()
}

func getBackwardS(s string) string {
	var backwardS strings.Builder

	for i := len(s) - 1; i >= 0; i-- {
		backwardS.WriteByte(s[i])
	}

	return backwardS.String()
}
