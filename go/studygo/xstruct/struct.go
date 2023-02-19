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

	x := &People{}
	x.arr = append(x.arr, 0)
	x.arr = append(x.arr, 1)
	x.arr = append(x.arr, 2)
	fmt.Println(x.arr)
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

func (x People) sayHello() {
	fmt.Println(x.name, " hello!")
}

func varFuncReceivePointer() {
	var p1 People
	p1.name = "xiaoming"
	p1p := &p1
	p1p.sayHello()
}

type XPeople struct {
	v [10]int
}

func (s XPeople) changev() {
	s.v[1] = 999
	fmt.Println(s)
}

func (s *XPeople) pchangev() {
	s.v[1] = 999
	fmt.Println(s)
}

func useVarAsReceiver() {
	var p XPeople = XPeople{}
	p.changev()
	fmt.Println(p)
}
func usePointerAsReceiver() {
	var p XPeople = XPeople{}
	p.pchangev()
	fmt.Println(p)

}
