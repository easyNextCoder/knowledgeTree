package xchannels

import (
	"fmt"
	"sync"
)

func chanChan() {
	chch := make(chan chan int)

	sw := sync.WaitGroup{}

	sw.Add(1)
	go func() {
		defer func() {
			sw.Done()
			fmt.Println("go1 exit")
		}()
		fmt.Println("go1")
		ch := make(chan int)
		chch <- ch
		ch <- 1
		fmt.Println("go1 通过ch发送的值", 1)

	}()

	sw.Add(1)
	go func() {
		defer func() {
			sw.Done()
			fmt.Println("go2 exit")
		}()
		fmt.Println("go2")
		ch := <-chch
		val := <-ch
		fmt.Println("go2 通过ch拿到的值", val)
	}()

	sw.Wait()
}
