package xchannels

import (
	"fmt"
	"time"
)

func cacheRead(ch chan int) {
	v := <-ch
	fmt.Println("cacheRead:", v)
}

func cacheWrite(ch chan int) {
	ch <- 0
	fmt.Println("cacheWrite ok")
}

func nullCacheReadDo() { //阻塞
	ch := make(chan int, 1)
	cacheRead(ch)

	time.Sleep(time.Second * 3)
}

func nullCacheWriteDo() { //不会阻塞
	ch := make(chan int, 1)
	cacheWrite(ch)
	time.Sleep(time.Second * 3)
}
