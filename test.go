package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 10; i++ {
		j := i
		go func() {
			fmt.Println(&i,&j)
			fmt.Println(i,j)
		}()
	}
	time.Sleep(1000000000)
	fmt.Println("over!!")
}
