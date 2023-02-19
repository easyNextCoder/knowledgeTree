package slice

import "fmt"

func slicePointerWork(ps *[]int) {
	s := ps
	for i := 0; i < 100; i++ {
		*s = append(*s, 1)
	}

}

func slicePointerWork2(ps *[]int) {
	x := []int{1, 2}
	slicePointerWork(&x)
	fmt.Println(x)
}
