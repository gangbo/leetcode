package main

import (
	"fmt"
	"reflect"
)

func main(){
	x,y := 1, 2
	var arr =  [...]int{5:2}
	//数组指针
	var pf *[6]int = &arr

	fmt.Println(reflect.TypeOf(pf))
	//指针数组
	pfArr := [...]*int{&x,&y}
	fmt.Println(pf)
	fmt.Println(pfArr)
}
