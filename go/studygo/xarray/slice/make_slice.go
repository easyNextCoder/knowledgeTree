package slice

import "fmt"

func makeSliceWork() {
	x := make([]int, 5, 10)
	x = append(x, 5)
	fmt.Println(x, len(x), cap(x))

	y := make([]int, 10)
	y = append(y, 5) //屏蔽前后cap不同
	fmt.Println(y, len(y), cap(y))
}
