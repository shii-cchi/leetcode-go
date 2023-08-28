package main

import "strings"

func wordPattern(pattern string, s string) bool {
	allWords := strings.Split(s, " ")

	if len(allWords) != len(pattern) {
		return false
	}

	matchMap := make(map[byte]string)
	backwardMap := make(map[string]byte)

	for i, _ := range pattern {
		patternChar := pattern[i]
		word := allWords[i]

		if value, ok := matchMap[patternChar]; ok {
			if value != word {
				return false
			}
		} else {
			if value, ok := backwardMap[word]; ok {
				if value != patternChar {
					return false
				}
			}

			matchMap[patternChar] = word
			backwardMap[word] = patternChar
		}
	}

	return true
}
