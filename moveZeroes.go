package main

import "fmt"

func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int)  {
	var zeroNotIn = make([]int, len(nums))
	var i = 0
	for _, v := range nums {
		if v == 0 {
			continue
		}
		zeroNotIn[i] = v
		i++
	}
	copy(nums, zeroNotIn)
}
