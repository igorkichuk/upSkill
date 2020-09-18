package main

import "fmt"

func main() {
	nums := []int{1, 6, 7, 8, 3, 4, 2, 13, 5, 15, 11, 10} //not sorted
	target := 14
	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	neededNums := map[int]int{} //map[need]index
	for i, num := range nums {
		neededIndex, ok := neededNums[num]
		if ok {
			return []int{i, neededIndex}
		}

		neededNums[target-num] = i
	}

	return []int{-1, -1}
}
