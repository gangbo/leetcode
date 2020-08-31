package main

import (
	"fmt"
	"math"
)

func main() {
	var x int = -121
	fmt.Printf("%t", isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return  false
	}
	ori := x
	var y int
	for x != 0 {
		y = y*10 + x%10
		x /= 10
	}
	if (y > math.MaxInt32) || (y < math.MinInt32) {
		return false
	}
	return ori == y
}

