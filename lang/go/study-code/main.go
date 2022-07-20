package main

import (
	_ "../study-code/testfile/unit/service"
	"fmt"
	"math"
	_ "net/http/pprof"
	"runtime"
	"time"
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

func main() {

	/*
		test use var define map
	*/

	vmp["name"] = "your"
	fmt.Println(vmp["name"])

	/*
		convey test
	*/

	//testFor()
	//testStruct()
	//testArray()
	//testRange()
	//testSlice()
	//testMap()
	//testSort()

	//go testGoruntines()
	//testGoruntines()

	//pp.Testwork()
	//testChannels()
	//testChannelsWithBuffer()

	//go goroutineLeak()
	//_ = http.ListenAndServe("0.0.0.0:6060", nil)
	//workOnChangeInterface()
}
