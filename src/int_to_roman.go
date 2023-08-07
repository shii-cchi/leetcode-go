//package main
//
//import (
//	"math"
//	"strconv"
//)
//
//func main() {
//
//}
//
//func intToRoman(num int) string {
//	romanNumerals := map[int]uint8{
//		1:    'I',
//		5:    'V',
//		10:   'X',
//		50:   'L',
//		100:  'C',
//		500:  'D',
//		1000: 'M',
//	}
//
//	splitedNumber := splitNumber(strconv.Itoa(num))
//}
//
//func splitNumber(s string) []int {
//	splitedNumber := make([]int, len(s))
//
//	for i, digit := range s {
//		splitedNumber = append(splitedNumber, int(digit)*int(math.Pow(10, float64(i))))
//	}
//
//	return splitedNumber
//}
