package xvar_definition

import (
	"fmt"
)

func newArr() {
	arr := new([]int) //arr存在，*arr也可用
	*arr = append(*arr, 0)
	fmt.Println(arr, *arr)
}

func newMap() {

	mp := new(map[int]bool)
	(*mp)[-1] = false //panic nil map
}

func newChan() {
	ch := new(chan int)
	fmt.Println(ch, *ch) //ch存在，*ch是nil
}

func newStruct() {
	//testTimer time.Timer
	//intVal    int
	//arrVal    []int
	//mapVal    map[int]bool
	//chanVal   chan int
	bs := new(Bus)
	fmt.Println(bs.testTimer.C, bs.intVal, bs.arrVal, bs.mapVal, bs.chanVal) //bs.testTimer.C是nil， bs.intVal存在，bs.arrVal存在，bs.mapVal是nil，bs.chanVal是nil
	bs.mapVal[-1] = true                                                     //nil map
}
