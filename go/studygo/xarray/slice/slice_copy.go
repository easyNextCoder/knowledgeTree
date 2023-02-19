package slice

import "fmt"

func SliceCopy() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	y := make([]int, 5)
	copy(y, x)
	fmt.Println("src: ", x, "des: ", y)
}
