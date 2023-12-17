package xpanic

import (
	"fmt"
	"time"
)

func directPanic() {
	panic("directPanic")
}

type Card int

func panicRecoverWrapper() int {

	//使用defer处理则正常返回
	go func() {

		defer func() { //panicRecover是在一个go程内，则此defer无法捕捉
			if e := recover(); e != nil {
				fmt.Println("outer panic recover err ", e)
			}
		}()

		go func() {
			res, b := panicRecover(map[int]bool{1: true})
			fmt.Println("panic return value:", res, b, len(res))
		}()

	}()

	fmt.Println("panicRecoverWrapper")

	time.Sleep(time.Second * 2)

	return 0
}

func panicRecover(a map[int]bool) (map[int]bool, error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panicRecover err(%s) a(%+v)\n", e, a)
		}
	}()

	func(a, b int) int {
		return a / b
	}(5, 0)

	fmt.Println("panicRecover")

	return map[int]bool{}, fmt.Errorf("this error")
}
