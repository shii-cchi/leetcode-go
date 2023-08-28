package main

import "strconv"

func fizzBuzz(n int) []string {
	var s []string

	for i := 0; i < n; i++ {
		if (i+1)%3 == 0 && (i+1)%5 == 0 {
			s = append(s, "FizzBuzz")
		} else if (i+1)%3 == 0 && (i+1)%5 != 0 {
			s = append(s, "Fizz")
		} else if (i+1)%3 != 0 && (i+1)%5 == 0 {
			s = append(s, "Buzz")
		} else {
			s = append(s, strconv.Itoa(i+1))
		}
	}

	return s
}
