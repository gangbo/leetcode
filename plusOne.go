package main

import (
	"fmt"
)

func main() {
	var testCases = [5]*[]int{
		{1, 2, 3, 5, 2},
		{9, 9, 9},
		{9},
		{2},
		{1, 9},
	}
	for _, v := range testCases {
		fmt.Println(">>>>>>>", *v)
		fmt.Println(plusOne(*v))
	}

}

func plusOne(digits []int) []int {
	var result []int
	var right []int
	for i := len(digits) - 1; i > -1; i-- {
		v := digits[i]
		if v < 9 {
			v++
			result = digits[0:i]
			right = append(right, v)
			break
		}
		right = append(right, 0)
		if i == 0 {
			right = append(right, 1)
		}
	}

	for i := len(right) - 1; i > -1; i-- {
		result = append(result, right[i])
	}

	return result
}
