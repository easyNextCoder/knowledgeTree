package xarray

import "fmt"

func Pic(dx, dy int) [][]uint8 {

	a := make([][]uint8, dy)
	for x := range a {
		b := make([]uint8, dx)
		for i := range b {
			b[i] = uint8(x*i - 1)
		}
		a[x] = b
	}
	return a
}

func testSlice() {
	res := Pic(5, 5)
	fmt.Println(res)
	for i := range res {
		for j := range res[i] {

			fmt.Print(res[i][j])
		}
	}

}
