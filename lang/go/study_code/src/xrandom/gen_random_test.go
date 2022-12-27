package xrandom

import (
	"fmt"
	"testing"
)

func Test_genRand(t *testing.T) {
	genRand()

	x := make([][]int, 4)
	for i := 0; i < 5; i++ {
		x = append(x, []int{2})
		x[i] = append(x[i], 1)
	}
	fmt.Println(x)
}
