package main

import "fmt"

func main() {
	//var nums = []int{2, 7, 11, 15}
	//var target = 9
	var nums = []int{3, 2, 4}
	var target = 6
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	result := make([]int, 0)
	for k1 := 0; k1 < len(nums); k1++ {
		for k2 := k1 + 1; k2 < len(nums); k2++ {
			if target == nums[k1]+nums[k2] {
				result = append(result, k1, k2)
				return result
			}

		}
	}
	return result
}
