package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func sum1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		time.Sleep(300 * time.Millisecond)
	}
	c <- sum
}

func testChannels() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	c := make(chan int)
	go sum(s[:len(s)/2], c)

	go sum1(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}

func testChannelsWithBuffer() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 10
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; ; i++ {
		c <- x
		x, y = y, x+y
		if x > 100 {
			close(c)
		}
	}
}

func testChannelsWithBufferClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func WriteN2Chan(ch chan int, n int) {
	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch)
	//chan不需要主动关闭，会被垃圾回收
	//只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。
	//发送方如果不close，会产生all goroutines asleep - deadlock!
	//close之后再往通道中写入，会panic:send on closed channel
}

func Write2Chan(ch chan int, val int) {
	ch <- val
}

func ReadFromChan(ch chan int) int {
	val := <-ch
	return val
}

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
