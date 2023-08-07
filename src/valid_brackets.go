package main

import "fmt"

func main() {
	s := "([[])"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	openBrackets := make([]rune, 0)
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			openBrackets = append(openBrackets, rune(s[i]))
		} else if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(openBrackets) == 0 {
				return false
			}

			if openBrackets[len(openBrackets)-1] != brackets[rune(s[i])] {
				return false
			}

			openBrackets = openBrackets[:len(openBrackets)-1]
		}
	}

	return len(openBrackets) == 0
}
