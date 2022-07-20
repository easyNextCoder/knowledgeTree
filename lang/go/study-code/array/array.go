package array

import (
	"fmt"
	"testing"
)

type arraySlice struct {
	ptr []int
	len int
	cap int
}

func (p *arraySlice) put(n int, val int) {
	if n < 0 || n >= p.len {
		panic("out of bounds")
	}
	p.ptr[n] = val
}

func (p *arraySlice) get(n int) int {
	if n < 0 || n >= p.len {
		panic("out of bounds")
	}
	return p.ptr[n]
}

func myLen(as arraySlice) int {
	return as.len
}

func myCap(as arraySlice) int {
	return as.cap
}

func myAppend(as arraySlice, value int) arraySlice {
	fmt.Println("work")
	if as.len+1 > as.cap {
		newLen := as.cap
		if newLen < 1 {
			newLen = 1
		} else {
			newLen = 2 * as.cap
		}

		ptrNew := make([]int, newLen)
		for i := 0; i < as.len; i++ {
			ptrNew[i] = as.ptr[i]
		}
		as.cap = newLen
		as.ptr = ptrNew
	}
	if as.len+1 > len(as.ptr) && as.len+1 <= cap(as.ptr) {
		as.ptr = as.ptr[:as.len+1]
	}
	as.ptr[as.len] = value
	as.len++
	return as
}

func mySlice(as arraySlice, l int, r int) arraySlice {
	if l < 0 || r > as.cap {
		panic("out of bounds")
	}
	as.ptr = as.ptr[l:r]

	as.len = r - l
	as.cap = cap(as.ptr)
	return as
}

func NewArray(n int) arraySlice {
	pt := make([]int, n)
	return arraySlice{
		ptr: pt,
		len: n,
		cap: n,
	}
}

func WorkOnNewArray(t *testing.T) {
	arr := NewArray(10)
	arr1 := make([]int, 10)
	if myLen(arr) != len(arr1) {
		t.Error("len not equal!")
	}
	if myCap(arr) != cap(arr1) {
		t.Error("cap not equal")
	}

	arr = mySlice(arr, 0, 1)
	arr1 = arr1[0:1]
	if myLen(arr) != len(arr1) {
		t.Error("slice(0, 1) len not equal!")
	}
	if myCap(arr) != cap(arr1) {
		t.Error("[0:1] cap not equal")
	}

	for i := 0; i < 100; i++ {
		arr = myAppend(arr, 1)
		arr1 = append(arr1, 1)
		fmt.Println(myLen(arr), len(arr1), myCap(arr), cap(arr1))
		if myLen(arr) != len(arr1) {
			fmt.Println("Error on: index ", i)
			t.Error("myAppend slice(0,1) len not equal!")
		}
		if myCap(arr) != cap(arr1) {
			t.Error("append [0:1] cap not equal")
		}
	}

}

func testChangeArr(arr []int) []int {
	for i, _ := range arr {
		arr[i] += 1
	}
	return arr
}

func useSliceAfterMake() {
	var TestGroups = make([]interface{}, 1)
	TestGroups = append(TestGroups, new(interface{}))
	fmt.Println(len(TestGroups), TestGroups)
}
