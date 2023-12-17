package xchannels

import (
	"fmt"
	"time"
)

func nilChannelWork() { //错误用法
	var x chan int //这里定义了一个nil的channel不可以使用

	go func() {
		fmt.Println("write start:")
		x <- 1 //在这里卡住
		fmt.Println("write to chan done")
	}()

	go func() {
		fmt.Println("read start:")
		val, ok := <-x //在这里卡住
		fmt.Println("read from chan done", ok, val)
	}()

	time.Sleep(time.Second * 2)
}

func notNilChannelWork() {
	var x chan int
	x = make(chan int)

	go func() {
		x <- 1 //panic
		fmt.Println("write to chan done")
	}()

	go func() {
		val, ok := <-x
		fmt.Println("read from chan done", ok, val)
	}()

	time.Sleep(time.Second * 2)

}
