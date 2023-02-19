package xselect

import (
	"fmt"
	"sync"
	"time"
)

func SendCommand(ch chan int) {
	ch <- 1
}

func Wait(ch chan int) {
	fmt.Println("run wait")
	select {
	case v := <-ch:
		fmt.Println("Wait ", v)
	}
	fmt.Println("run here")
}

func blockWaitWork() {
	ch := make(chan int, 0)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		Wait(ch)
	}()
	time.Sleep(time.Second * 2)
	wg.Add(1)
	go func() {
		defer wg.Done()
		SendCommand(ch)
	}()
	wg.Wait()
}
