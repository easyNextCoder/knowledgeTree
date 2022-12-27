package xcopy

import (
	"fmt"
)

type req struct {
	ival int
	sval string
	arr  []int
}

//浅拷贝,改变一个数组中的内容，另一个数组的内容也发生改变
func shadowCopy() {
	req1 := new(req)
	req1.ival = 10
	req1.sval = "work"
	req1.arr = make([]int, 0)
	req1.arr = append(req1.arr, 1)
	req1.arr = append(req1.arr, 2)
	req1.arr = append(req1.arr, 3)

	req2 := new(req)

	a := make([]*req, 2)
	b := []*req{req1, req2}
	copy(a, b)
	a[0].ival = 100

	fmt.Println(a, b, b[0].ival)
	//
	//req2 := new(req)
	//*req2 = *req1 //import
	//req2.arr[1] = 9
	//req2.ival = 89
	//fmt.Printf("%d, %d, %+v\n", unsafe.Pointer(&req1.arr), unsafe.Pointer(&req2.arr), req1)
	//fmt.Println(req1, *req1, req2, *req2)

}

//深拷贝，改变一个数组的内容，另一个数组的内容不变
func deepCopy() {
	a := []string{"a", "b", "c"}
	b := make([]string, len(a), cap(a))
	copy(b, a) //import
	a[1] = "99"
	fmt.Println(a, b)
}
