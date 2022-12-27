package xfloat

import "fmt"

func upperDevide(a, b int) int {
	integerRes := int(float64(a) / float64(b))
	if integerRes*b < a {
		return integerRes + 1
	}
	return integerRes
}

const (
	a = iota
	b
	c
	d
	e = 7
	f
	g
)

func run() {
	//a := 5
	//a := make(map[int]int)
	//fmt.Println("a10 is:", a[10])
	//fmt.Println(upperDevide(15, 4))
	//
	//x := "work"
	//fmt.Println(x[0])

	fmt.Println(a, b, c, d, e, f, g)

}
