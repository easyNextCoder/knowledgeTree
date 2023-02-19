package slice

import "fmt"

func nilSliceWork() {
	var x []int
	x = append(x, 1)
	fmt.Println("var 申请", x[0]) //ok

	y := []int{}
	y = append(y, 1)
	fmt.Println(":= 申请", y[0]) //ok
}
