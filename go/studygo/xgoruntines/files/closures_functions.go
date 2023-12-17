package files

import (
	"fmt"
	"time"
)

var pts [3][3]int = [3][3]int{
	[3]int{1, 2, 3},
	[3]int{3, 2, 1},
	[3]int{10, 20, 30},
}

func closuresWork() {
	for i := 0; i < 3; i++ {
		xv := &pts[i]
		f := func() {
			fmt.Println("now result is", xv)
		}
		time.AfterFunc(time.Second*time.Duration(i+1), f)
	}
	fmt.Println("register done")

	time.Sleep(time.Second * 10)

}

func inner(a int) {
	time.Sleep(time.Second * 1)
	fmt.Println(a)
}

func outerInner() {
	mp := map[int]int{
		1: 1,
		2: 2,
	}
	for k, _ := range mp {
		go inner(k)
	}

	mp = make(map[int]int)
	fmt.Println(mp)
	time.Sleep(time.Second * 5)

}
