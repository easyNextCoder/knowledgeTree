package xvar_definition

import "fmt"

//make([]interface{}, 1)

func make0() {
	var arr0 = make([]interface{}, 0)
	arr0 = append(arr0, new(interface{}))
	fmt.Println("make0 arr0:", arr0)
}

func make1() {
	var arr1 = make([]interface{}, 1)
	arr1 = append(arr1, new(interface{}))
	fmt.Println("make1 arr1:", arr1)
}
