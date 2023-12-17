package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	func(v, m int) {
		fmt.Println(v, math.MaxInt32, v > math.MaxInt32, v*2500, unsafe.Sizeof(v), unsafe.Sizeof(m), unsafe.Sizeof(1), unsafe.Sizeof(bool(true)))
	}(int(9223372036854775807), 10)
}
