package xbit

import "fmt"

//16进制的位移动操作，只用了后两位
func CalSource(oldStatus, sourceVal int) int {
	last := oldStatus & 15
	res := (last << 4) | sourceVal
	return res
}

func check(x int64, index int) {
	x = 17
	x = x & (1 << 10)
	a := int32(-1)
	for i := 0; i < 32; i++ {
		fmt.Println(a & (1 << i))
	}

	fmt.Printf("%064b, %b, %064b", x, int32(2147483647), (int64(-1)))
	y := (2 | 1<<0)
	fmt.Println("Y is:", y)
}

func runCheck() {
	check(0, 0)
}
