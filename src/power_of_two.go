package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(512))
}

func isPowerOfTwo(n int) bool {
	if n == 0 || (n%2 != 0 && n != 1) {
		return false
	}

	for n != 1 {
		if n%2 != 0 {
			return false
		}

		n = n / 2
	}

	return true
}
