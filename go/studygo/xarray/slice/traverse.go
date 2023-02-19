package slice

import "fmt"

type Traverse struct {
	v     [3]int
	b     int
	slice []int
}

func changeRangeItem() {
	var x Traverse = Traverse{
		v:     [3]int{1, 2, 3},
		b:     100,
		slice: []int{9, 8, 7},
	}

	var y Traverse = Traverse{
		v:     [3]int{1, 2, 3},
		b:     100,
		slice: []int{9, 8, 7},
	}

	var s []Traverse = []Traverse{x, y}
	//这样定义变量是以复制的方式进行的，
	//对于slice类型的变量复制了SliceHeader,所以只有改变slice中的内容会改变原被append的变量的内容，均不会改变

	fmt.Printf("变量放入数组之后和之前的地址比较\n x(%p) x.v(%p) x.b(%p) x.slice(%p%p) \n"+
		"s[0](%p) s[0].v(%p) s[0].b(%p) s[0].slice(%p%p) \n"+
		"y(%p) y.v(%p) y.b(%p) y.slice(%p%p)\n"+
		"s[1](%p) s[1].v(%p) s[1].p(%p) s[1].slice(%p%p)\n",
		&x, &x.v, &x.b, x.slice, &x.slice,
		&s[0], &s[0].v, &s[0].b, s[0].slice, &s[0].slice,
		&y, &y.v, &y.b, y.slice, &y.slice,
		&s[1], &s[1].v, &s[1].b, s[1].slice, &s[1].slice)

	var sv []Traverse

	for _, item := range s {
		//每次遍历item会复制s数组中的内容，且item变量会一直被复用
		//item := item //这样写会禁止复用
		item.v[1] = 999     //无法改变s中的值
		item.b = 999        //无法改变s中的值
		item.slice[1] = 999 //可以改变s中的slice的值，但无法改变slice的len和cap
		item.slice = append(item.slice, 999)
		sv = append(sv, item)
		fmt.Printf("遍历时候的中间变量的地址是:%p %v\n", &item, item)
	}

	fmt.Printf("变量放入数组之后和之前的地址比较\n x(%p) x.v(%p) x.b(%p) x.slice(%p%p)\n"+
		"sv[0](%p) sv[0].v(%p) sv[0].b(%p) sv[0].slice(%p%p)\n"+
		"y(%p) y.v(%p) y.b(%p) y.slice(%p%p)\n"+
		"sv[1](%p) sv[1].v(%p) sv[1].p(%p) sv[1].slice(%p%p)\n",
		&x, &x.v, &x.b, x.slice, &x.slice,
		&sv[0], &sv[0].v, &sv[0].b, sv[0].slice, &sv[0].slice,
		&y, &y.v, &y.b, y.slice, &y.slice,
		&sv[1], &sv[1].v, &sv[1].b, sv[1].slice, &sv[1].slice)

	fmt.Println(s)
	fmt.Println(sv)

}
