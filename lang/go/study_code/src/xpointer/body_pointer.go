package xpointer

import (
	"fmt"
	"time"
)

type x struct {
	j int
	k string
}

func pointer_pointer() {
	var costTime time.Duration = time.Duration(0)
	pointer_pointer1(&costTime)
	pointer_pointer2(&costTime)
	fmt.Println(costTime)
}

func pointer_pointer1(pCostTime *time.Duration) {
	fmt.Println(pCostTime, *pCostTime)
	*pCostTime += time.Millisecond * 16 * 1000

}

func pointer_pointer2(pCostTime *time.Duration) {
	fmt.Println("p2", pCostTime, *pCostTime)
	*pCostTime += time.Millisecond * 1 * 1000
	if time.Second*20-*pCostTime < time.Second*5 {
		fmt.Println("时间不够了赶紧做！")
	}
	*pCostTime += time.Second * 2
	a := *pCostTime
	fmt.Println("a", a.Milliseconds())
}

func work1() {

	v := new(x)
	v.k = "hello"
	v.j = 10
	p := *v
	p.j = 20
	p.k = "world"

	a := []int{1, 2, 3, 4}
	b := []int{5, 6}
	a = append(a, b...)
	b = b[:0]

	fmt.Println(a, b)
}

func run() {
	pointer_pointer()

}
