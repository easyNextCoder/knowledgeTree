package xchannels

import (
	"fmt"
	"time"
)

func ReadWriteUninitializedChan() {
	var ch chan int
	if ch == nil {
		fmt.Println("read write uninitializedChan nil")
	}
	go Write2Chan(ch, 1)
	ret := <-ch
	fmt.Println("ret is:", ret)
	//nil管道接收和发送都会出现阻塞
}

func ReadWriteInitializedChan() {
	ch := make(chan int)
	go Write2Chan(ch, 1)
	ret := <-ch
	fmt.Println("ret is:", ret)
	//无缓冲的chan，单独只发送或者接收都会出现阻塞
}

func ReadWriteInitializedChanWithBuffer() {
	ch := make(chan int, 15)
	go WriteN2Chan(ch, 10)
	//如果使用goroutine去写入，超过buffer容量会自动退出，不会报错

	for {
		//如果发送方不主动close，接收方会报错
		//如果发送方主动close，ok变量会为false
		fmt.Println("wait to read.")
		if data, ok := <-ch; ok {
			fmt.Println("ok", data)
		} else {
			fmt.Println("读空的已经关闭的chan", <-ch, <-ch, <-ch)
			break
		}
	}
	time.Sleep(200 * time.Millisecond)
}
