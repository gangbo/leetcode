package main

import (
	"fmt"
	"math"
)

func main() {
	//x := -123
		//var x int = -2147483648
	var x int = -21474
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	var y int
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}
	if (y > math.MaxInt32) || (y < math.MinInt32) {
		return 0
	}
	return y
}
