package main

import "fmt"

func main() {
	nums := []int{1, 2, 2}
	target := 4
	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	var result []int
	found := false

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	return result
}
