package main

import (
	"fmt"
	"strings"
)

func main() {
	var strs = []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if 0 == len(strs) {
		return ""
	}
	var prefix = strs[0]
	for _, v := range strs {

		for 0 != strings.Index(v, prefix) {
			prefix = prefix[:len(prefix)-1]
		}

	}
	return prefix
}
