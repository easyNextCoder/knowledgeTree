package files

import (
	"fmt"
	"time"
)

func goes(i int) {
	fmt.Println(i)
}

func goInForRange() {
	fmt.Println("使用普通函数")
	x := []int{1, 2, 3}
	for _, v := range x {
		go goes(v)

	}
	time.Sleep(time.Second * 1)
}

func goFuncInForRange() {
	fmt.Println("使用闭包函数")
	x := []int{1, 2, 3}
	for _, v := range x {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(time.Second * 1)
}
