package xchannels

import (
	"fmt"
	"sync"
	"time"
)

type (
	inner struct {
		i  [10]int
		ip interface{}
	}
	outer struct {
		inn inner
		i   int
	}
)

func structPointer() {
	schan := make(chan *outer)
	cchan := make(chan struct{})

	iou := outer{
		inn: inner{i: [10]int{1, 2, 3, 4}, ip: schan},
		i:   10,
	}

	sw := sync.WaitGroup{}
	sw.Add(1)
	go func() {
		defer sw.Done()
		fmt.Printf("发送的结构体(%v) 结构体指针(%p)\n", iou, &iou)
		schan <- &iou
		<-cchan

		fmt.Printf("改动过的结构体(%v) 结构体指针(%p)\n", iou, &iou)

	}()

	sw.Add(1)
	go func() {
		defer sw.Done()
		if val, ok := <-schan; ok {
			fmt.Printf("收到的结构体(%v) 结构体指针(%p)\n", val, &val)
			val.inn.i = [10]int{9, 9, 9, 9, 9, 9}
			time.Sleep(time.Second * 3)
			cchan <- struct{}{}
		}

	}()

	sw.Wait()
}

func channelStructWork() {
	schan := make(chan outer)
	cchan := make(chan struct{})

	iou := outer{
		inn: inner{i: [10]int{1, 2, 3, 4}, ip: schan},
		i:   10,
	}

	sw := sync.WaitGroup{}
	sw.Add(1)
	go func() {
		defer sw.Done()
		fmt.Printf("发送的结构体(%v) 结构体指针(%p)\n", iou, &iou)
		schan <- iou
		<-cchan

		fmt.Printf("改动过的结构体(%v) 结构体指针(%p)\n", iou, &iou)

	}()

	sw.Add(1)
	go func() {
		defer sw.Done()
		if val, ok := <-schan; ok {
			fmt.Printf("收到的结构体(%v) 结构体指针(%p)\n", val, &val)
			val.inn.i = [10]int{9, 9, 9, 9, 9, 9}
			time.Sleep(time.Second * 3)
			cchan <- struct{}{}
		}

	}()

	sw.Wait()

}
