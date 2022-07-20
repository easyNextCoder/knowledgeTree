package goruntines

import (
	"fmt"
	"time"
)

func work(name string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(name)
	}
}

func testGoruntines() {
	go work("world")
	work("hello")

}

/*
	1. go什么时间结束协程？
		go在主函数结束的时候结束协程
	2. 当遇到崩溃的时候应当做什么处理？
		使用defer+recover来捕获panic，并进行处理
	3. 课后拓展：什么为go泄漏，怎么定位哪些go泄漏了(pprof)
		go101
		如何定位泄漏
*/

func PanicThenRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("PanicThenRecover 函数发生错误")
		}
	}()
	fmt.Println("PanicThenRecover start")
	var mp map[string]int
	//mp["hello"] = 1
	fmt.Println("WORK", mp["work"])
}

func GoPanicThenRecover() {
	go PanicThenRecover()
	go PanicThenRecover()
	fmt.Println("主函数运行正常！")
	time.Sleep(200 * time.Millisecond)
}
