package main

import "fmt"

func main() {
	nums := []int{-5, -3, -1, 2, 3, 5, 6, 9, 12, 15} // sorted
	target := 14
	result := twoSum2(nums, target)
	fmt.Println(result)
}

func twoSum2(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for left != right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum > target {
			right--
		}
		if sum < target {
			left++
		}
	}

	return []int{}
}
