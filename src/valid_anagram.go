package main

import "fmt"

func main() {
	
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make(map[rune]int)
	mapT := make(map[rune]int)

	for _, value := range s {
		mapS[value]++
	}

	for _, value := range t {
		mapT[value]++
	}

	if len(mapS) != len(mapT) {
		return false
	}

	return fmt.Sprint(mapS) == fmt.Sprint(mapT)
}
