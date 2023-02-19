package xmap

import "fmt"

type Proj struct {
	p1 int
	a1 []int
}

type mpAsFuncParamStruct struct {
	mp  map[int]int
	val *Proj
}

func mapAsFuncParam(x *mpAsFuncParamStruct) {
	mp := x.mp
	mp[1] = 1
	mp[2] = 2
	mp[3] = 3
	pp := x.val
	pp.p1 = 1000
}

func mapAsFuncParam1() {
	mp := map[int]int{}
	x := 10

	p := &mpAsFuncParamStruct{
		mp: mp,
		val: &Proj{
			p1: -1,
			a1: []int{1, 2, 3},
		},
	}
	fmt.Println(mp, x, *p.val)
	mapAsFuncParam(p)
	fmt.Println(mp, x, *p.val)
}
