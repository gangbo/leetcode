package main

import "fmt"

func main() {
	var strs = []string{"()", "()[]{}", "(]"}

	for _, v := range strs {
		fmt.Println(v, isValid(v))
	}

}

/**
https://leetcode-cn.com/problems/valid-parentheses/
*/

func isValid(s string) bool {
	fmt.Println(s[0:1])
	var stack []string
	for i:=0; i<len(s); i++ {

	}
	append(stack)
	return true
}
