package xbuildin_funcs

import "fmt"

type People struct {
	weight int
	height int
	name   string
	arr    []int
}

func copyStructLen0Arr() {
	a := People{
		weight: 0,
		height: 0,
		name:   "",
		arr:    []int{1, 2, 3},
	}
	b := new(People)
	b.arr = make([]int, len(a.arr)) //必须传入长度才能正常拷贝
	a.arr = append(a.arr, 0)
	a.arr = append(a.arr, 1)
	a.arr = append(a.arr, 2)
	copy(b.arr, a.arr)
	fmt.Println(a, b)
}
