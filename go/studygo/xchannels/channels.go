package xchannels

import (
	"fmt"
	"testing"
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

func testChannelsWithBuffer(t testing.T) {
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

func ReadFromChan(ch chan int) int {
	val := <-ch
	return val
}

type xs struct {
	v  [10]int
	vp []int
}

func chanSendStruct() {
	ch := make(chan interface{})

	go func() {
		fmt.Println("into go")
		v := <-ch
		fmt.Println(v)
	}()

	time.Sleep(5 * time.Second)
	ch <- xs{}

}
