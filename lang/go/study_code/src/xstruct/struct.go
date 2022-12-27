package xstruct

import (
	"fmt"
	"unsafe"
)

type Vertex struct {
	X int
	Y int
}

type People struct {
	weight int
	height int
	name   string
	arr    []int
	x      *Vertex
}

func testVar() {
	var a *People
	if a == nil {
		fmt.Print("a is nil")
		fmt.Println(-2 % 4)
	}
}

func testStruct() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println("testStruct:")
	fmt.Println(v.X)
	//
	//xyk := People{10, 10, "徐永康"}
	//fmt.Println(xyk.name, xyk.height, xyk.weight)

	var p1 People
	p2 := new(People)
	fmt.Println("新的测试：", unsafe.Sizeof(p1), unsafe.Sizeof(p2), unsafe.Sizeof(*p2), p1.name, p2.name)
}
