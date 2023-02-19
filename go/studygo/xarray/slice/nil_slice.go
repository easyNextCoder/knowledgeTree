package slice

import (
	"fmt"
	"reflect"
	"unsafe"
)

func resetVarSliceToNil() {
	var ssa [][]int
	var sa []int
	sa = append(sa, 1)
	ssa = append(ssa, sa)
	sa = nil
	fmt.Printf("now sa is: %v %p pp%p\n", sa, sa, &sa)
	sah := (*reflect.SliceHeader)(unsafe.Pointer(&sa))
	fmt.Printf("sliceHeader %x, %d, %d ", sah.Data, sah.Len, sah.Cap)
	sa = append(sa, 10)
	ssa = append(ssa, sa)

	fmt.Println(ssa)

}
