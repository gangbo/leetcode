package main

import "fmt"

func main() {

	var (
		a  [2]*int
		ia int
		ib int
	)
	a = [2]*int{&ia, &ib}
	fmt.Println(a)
	fmt.Println("array=>>>>>>>>>>")
	var b [3]*[]int = [3]*[]int{
		{2,3},
		{2,3,8},
		{2,3},
	}
	//b[0] = &[]int{2, 3}
	//b[1] = &[]int{4, 3}
	//b[2] = &[]int{2, 3}
	fmt.Println("bbbb=>>>>>>>>>>>>")
	//fmt.Println(b)
	for k, v := range b {
		fmt.Println(v, &b[k])
	}

	fmt.Println(b)
}
