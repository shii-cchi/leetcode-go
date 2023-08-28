package main

import "bytes"

func decodeMessage(key string, message string) string {
	substitutionMap := make(map[rune]byte)
	substitutionMap[' '] = ' '

	var decodedMes bytes.Buffer

	for _, value := range key {
		if _, ok := substitutionMap[value]; !ok {
			substitutionMap[value] = byte(97 + len(substitutionMap) - 1)
		}
	}

	for _, value := range message {
		decodedMes.WriteByte(substitutionMap[value])
	}

	return decodedMes.String()
}
