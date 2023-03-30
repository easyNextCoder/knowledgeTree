package xvar_definition

import "fmt"

//make 只是用来定义chan slice map 并且生成的不是指针
func make0() {
	var arr0 = make([]interface{}, 0)
	arr0 = append(arr0, new(interface{}))
	v := new([]int)
	*v = append(*v, 1)
	fmt.Println("make0 arr0:", arr0, *v)
}

func make1() {
	var arr1 = make([]interface{}, 1)
	arr1 = append(arr1, new(interface{}))
	fmt.Println("make1 arr1:", arr1)

}
