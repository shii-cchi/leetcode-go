package main

import "fmt"

func main() {
	s := "MCMXCIV"

	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	resultNumber := 0

	for i, symbol := range s {
		if i != len(s)-1 && needSubstract(rune(symbol), rune(s[i+1])) {
			resultNumber -= romanNumerals[rune(symbol)]
		} else {
			resultNumber += romanNumerals[rune(symbol)]
		}
	}

	return resultNumber
}

func needSubstract(symbol, nextSymbol rune) bool {
	return (symbol == 'I' && (nextSymbol == 'V' || nextSymbol == 'X')) || (symbol == 'X' && (nextSymbol == 'L' || nextSymbol == 'C')) ||
		(symbol == 'C' && (nextSymbol == 'D' || nextSymbol == 'M'))
}
