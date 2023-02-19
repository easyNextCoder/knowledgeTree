package slice

import "fmt"

type StructSlice struct {
	arr []int
	s   string
}

func StructSliceWork(p *StructSlice) {
	p.arr = append(p.arr, 1)
	p.arr = append(p.arr, 1)
	p.arr = append(p.arr, 1)
	p.arr = append(p.arr, 1)
	p.arr = append(p.arr, 1)
	fmt.Println("StructSliceWork", len(p.arr), cap(p.arr), p)
}

func StructSliceWorkWrapper() {

	valp := &StructSlice{
		arr: []int{1, 2, 3, 4},
	}

	fmt.Println("StructSliceWorkWrapper", len(valp.arr), cap(valp.arr), valp, &valp.arr)
	fmt.Printf("%p", &valp.arr)

	StructSliceWork(valp)

	fmt.Println("StructSliceWorkWrapper", len(valp.arr), cap(valp.arr), valp, &valp.arr)
	fmt.Printf("%p", &valp.arr)
}

type StructSliceP struct {
	arr *[]int
}

func StructSlicePWork(p *StructSliceP) {
	*p.arr = append(*p.arr, 1)
	*p.arr = append(*p.arr, 1)
	*p.arr = append(*p.arr, 1)
	*p.arr = append(*p.arr, 1)
	*p.arr = append(*p.arr, 1)
	fmt.Println("StructSlicePWork", len(*p.arr), cap(*p.arr), p)
}

func StructSlicePWorkWrapper() {

	valp := &StructSliceP{
		arr: &[]int{1, 2, 3, 4},
	}

	fmt.Println("StructSlicePWorkWrapper", len(*valp.arr), cap(*valp.arr), valp)

	StructSlicePWork(valp)

	fmt.Println("StructSlicePWorkWrapper", len(*valp.arr), cap(*valp.arr), valp)

}

func changeInForRange() {
	var x StructSlice = StructSlice{
		arr: []int{1, 2, 3, 4, 5},
		s:   "yourname",
	}
	var y StructSlice = StructSlice{
		arr: []int{1, 2, 3, 4, 5},
		s:   "myname",
	}

	var vl []StructSlice = []StructSlice{x, y}
	for _, p := range vl {
		p.s = "xxxxname"              //不会改变
		p.arr[1] = 100                //会发生改变
		p.arr = append(p.arr, 999)    //不会append成功，只对临时变量生效了
		fmt.Printf("临时变量的地址%p\n", &p) //for中迭代的变量只有一个会一只复用
	}
	fmt.Println(vl[0], vl[1])

	var sl []*StructSlice = []*StructSlice{&x, &y}
	for _, p := range sl {
		p.s = "xxxxname"           //会发生改变
		p.arr[1] = 100             //会发生改变
		p.arr = append(p.arr, 999) //会append成功，直接通过临时变量操作了原变量
	}
	fmt.Println(sl[0], sl[1])
}
