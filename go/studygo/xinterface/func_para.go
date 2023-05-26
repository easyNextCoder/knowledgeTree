package xinterface

import (
	"fmt"
)

type (
	s struct {
		i   int
		arr [10]int
	}
)

func interfaceWork(v interface{}) {
	fmt.Println("interfaceWork:", v)

	//使用v.(type)判断输入变量的类型
	switch v.(type) {
	case int:
		fmt.Println(v)
	case s:
		fmt.Println(v, "struct s")
	case *s:
		fmt.Println(v, "struct *s")
	default:
		fmt.Println("default")
		fmt.Println("default")
		fmt.Println("default")
	}

}

func interfaceWorkWrapper() {
	sv := s{}
	interfaceWork(sv)
	interfaceWork(&sv)
	va := []int64{}
	va = append(va, 1)
	fmt.Println(va)

	//var x interface{}
	//xi := x.(int)
	//fmt.Println(xi)
}

//only for test
