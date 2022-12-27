package xalg

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

func test() [][]int {
	fmt.Println("do it")
	ret := make([][]int, 2)
	for i, _ := range ret {
		ret[i] = append(ret[i], 0)
		ret[i] = append(ret[i], 1)
		ret[i] = append(ret[i], 2)
	}
	return ret
}

var xt = test()

func worktest2() {
	fmt.Println(xt)
	fmt.Println(xt)
	fmt.Println(xt)
}

func WorkTest() {
	fmt.Println(xt)
	worktest2()
	a := []int{1, 2, 3, 4}
	b := a[0:4]
	fmt.Println("the b res")
	fmt.Println(b)
}

func SortSlice() {
	arrLen := 10
	arr := make([]*person, 10)
	for i := 0; i < arrLen; i++ {
		arr[i] = &person{
			name: "xyk",
			age:  arrLen - i,
		}
	}
	fmt.Println(arr[0])
	func() {
		sort.Slice(arr, func(l, r int) bool {
			return arr[l].age < arr[r].age
		})
	}()

	fmt.Println(arr[0])

}
