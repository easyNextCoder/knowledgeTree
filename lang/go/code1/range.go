package main

import "fmt"

//range可以循环遍历切片或者映射
var pow = []int{-1, -2, -3, -4, -5, -6}

func testRange() {
	for index, value := range pow {
		fmt.Println(index, value)
	}
}
