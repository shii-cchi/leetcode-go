package main

import "fmt"

func main() {
	fmt.Println(mySqrt(1))
}

func mySqrt(x int) int {
	left := 0
	var right int

	if x > 1 {
		right = x
	} else {
		right = 1
	}

	middle := (left + right) / 2

	if middle == 0 {
		middle = 1
	}

	for middle > left {
		if middle*middle > x {
			right = middle
		} else {
			left = middle
		}
		middle = (left + right) / 2
	}

	return middle
}
