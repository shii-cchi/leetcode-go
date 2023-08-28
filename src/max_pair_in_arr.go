package main

func maxSum(nums []int) int {
	sumMap := make(map[int]int)
	maxSumNum := -1

	for _, num := range nums {
		if value, ok := sumMap[getMaxDigits(num)]; !ok {
			sumMap[getMaxDigits(num)] = num
		} else {
			if value < num {
				sumMap[getMaxDigits(num)] = num
			}

			if value+num > maxSumNum {
				maxSumNum = value + num
			}
		}
	}

	return maxSumNum
}

func getMaxDigits(num int) int {
	var max int

	for num > 0 {
		digit := num % 10
		if max < digit {
			max = digit
		}
		num /= 10
	}

	return max
}
