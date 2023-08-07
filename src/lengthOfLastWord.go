package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "the moon  "
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	arrWords := strings.Split(strings.TrimSpace(s), " ")
	return len(arrWords[len(arrWords)-1])
}
