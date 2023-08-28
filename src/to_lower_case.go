package main

import "strings"

func main() {

}

func toLowerCase(s string) string {
	var letters strings.Builder

	for _, value := range s {
		if value >= 65 && value <= 90 {
			value += 32
		}
		letters.WriteRune(value)
	}

	return letters.String()
}
