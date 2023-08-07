package main

import "fmt"

func main() {
	fmt.Println(addDigits(0))
}

func addDigits(num int) int {
	if num == 0 {
		return 0
	}

	digit := 0
	for num != 0 {
		digit += num % 10
		num = num / 10

		if num == 0 && digit > 9 {
			num = digit
			digit = 0
		}
	}

	return digit
}
