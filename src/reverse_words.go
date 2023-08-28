package main

import "strings"

func reverseWords(s string) string {
	allWords := strings.Split(s, " ")

	var builderS strings.Builder

	for i, str := range allWords {
		for i := len(str) - 1; i >= 0; i-- {
			builderS.WriteByte(str[i])
		}
		if i != len(allWords)-1 {
			builderS.WriteByte(' ')
		}
	}

	return builderS.String()
}
