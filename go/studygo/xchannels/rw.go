package xchannels

import (
	"fmt"
	"time"
)

func closeAfterWrite(ch chan int, n int) {
	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch)
	//chan不需要主动关闭，会被垃圾回收
	//只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。
	//发送方如果不close，会产生all goroutines asleep - deadlock!

	//往closed的通道写入: panic:send on closed channel
	//从closed的通道读取: 正常返回0
}

func ReadWrite() {
	ch := make(chan int)
	go closeAfterWrite(ch, 100)

	forLoopRead(ch)

	//go forRangeRead(ch)

	time.Sleep(4000 * time.Millisecond)
}

func forLoopRead(ch chan int) {
	for {
		//如果发送方不主动close，接收方会 deadlock(用go程跑不会panic但是内存泄漏了)
		//如果发送方主动close，ok变量会为false
		data, ok := <-ch
		if ok {
			fmt.Println("loop read", data)
		} else {
			fmt.Println("loop read 读空的已经关闭的chan", <-ch, <-ch, <-ch)
			//break
		}
	}
}

func forRangeRead(ch chan int) {
	for val := range ch {
		fmt.Println("range read", val)
	}
}
