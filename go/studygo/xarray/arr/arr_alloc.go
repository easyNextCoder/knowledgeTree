package arr

import "fmt"

type X struct {
	int
	arr  [10]int
	arrp []int
}

var xarr [4]*X

func load() {
	for i := range xarr {
		xarr[i] = new(X)
	}
}

func arrAllocWork() {
	load()
	for i, v := range xarr {
		v.arrp = append(v.arrp, i)
		fmt.Printf("this is %d %+v\n", i, v)
	}
}
