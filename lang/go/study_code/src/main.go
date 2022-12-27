package main

import (
	"fmt"
	"math"
	_ "net/http/pprof"
	"runtime"
	"time"
	xnsq "xnsq"
)

func add(a int, b int) int {
	return a + b
}

func swap(a string, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func testMain() {
	fmt.Println(add(1, 2))
	fmt.Println(swap("first", "second"))
	fmt.Println(split(10))

	var c, python, java bool
	fmt.Println(c, python, java)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint32 = uint32(f)
	fmt.Println(x, y, z)
}

type chanVal struct {
	a int
}

var ch = make(chan chanVal, 10)

func goroutineLeak() {
	defer func() {
		fmt.Println("groutines: ", runtime.NumGoroutine())

	}()

	go goroutineCollect()
	for i := 0; i < 20; i++ {
		go func() {
			ch <- chanVal{a: i}
		}()
		time.Sleep(time.Millisecond * 1000)
	}

}

func goroutineCollect() {
	defer func() {
		fmt.Println("goroutineCollect execute done")
	}()
	for i := 0; i < 5; i++ {
		go func() {
			rv := <-ch
			fmt.Println("collecting goroutine ", rv.a)
		}()
	}
}

var vmp = make(map[string]string)

func ect() {
	v := 0
	v++
	t := v
	t++
}

type XX struct {
	testTimer <-chan time.Time
	//testTimer time.Timer
}

func main() {

	//ysvc.UseYorm()
	//ysort.WorkTest()
	//xmap.Safe_map()
	//mp := make([]int, 2)
	//mp[0] = 4
	//mp[1] = 3
	//for v := range mp {
	//	fmt.Println(v)
	//}
	//fmt.Println(time.Now())
	//for i := 0; i < 8000000; i++ {
	//	ect()
	//}
	//fmt.Println(time.Now())

	//hand.SimpleWork()
	//tMap.UseValueAfterDelete()
	//tArray.OutTestReturnNil()
	//tMap.SimpleInitMap()

	//xx := new(XX)
	//xx.testTimer = time.After(time.Second * 2)
	//yy := time.NewTimer(time.Second * 2)
	//xx.testTimer = *yy
	//xx2 := new(XX)
	//xx2.testTimer = time.After(time.Second * 5)
	//yy := time.NewTimer(time.Second * 5)
	//&xx.testTimer = yy
	//xx.testTimer = time.After(time.Second * 3)
	//xx2 := new(XX)
	//go func() {
	//	for {
	//		select {
	//		case <-xx2.testTimer:
	//			fmt.Println("xxxxxhhhhhhhhhhhhhxxx")
	//		case <-xx.testTimer:
	//			//fmt.Println(xx.testTimer, xx.testTimer.C)
	//
	//			fmt.Println(xx.testTimer, xx.testTimer)
	//			fmt.Println("Timer off xxxxxxxxxxxxxxxxxxxxx")
	//			xx2.testTimer = time.After(time.Second * 5)
	//		}
	//	}
	//}()
	//time.Sleep(time.Second * 4)
	//fmt.Println("xyk")
	//xx.testTimer = time.After(2)

	//xx2.testTimer = time.After(time.Second * 4)
	//xx.testTimer = time.After(time.Second * 3)

	//fmt.Println(yy.Stop())
	//res := false
	//fmt.Println(res, xx.testTimer, xx.testTimer.C)
	//var s string
	//fmt.Scanln(&s)

	xnsq.StartNsq()

}
