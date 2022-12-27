package xchannels

import (
	"fmt"
	"sync"
	"time"
)

func chanChan() {
	chch := make(chan chan int)

	sw := sync.WaitGroup{}

	sw.Add(1)
	go func() {
		defer func() {
			sw.Done()
			fmt.Println("go2 exit")
		}()
		ch := make(chan int)
		fmt.Println("go1")
		chch <- ch
		fmt.Println("go1 通过chch发送ch", chch)
		//ch <- 1
		//fmt.Println("go1 通过ch发送的值", 1)

		select {
		//阻塞读
		case ch <- 1:
			fmt.Println("尝试往内部ch中写入")
			//default:
			//	fmt.Println("尝试往内部ch中写入，失败")
		}

		fmt.Println("go1 通过ch发送的值", 1)

	}()

	sw.Add(1)
	go func() {
		defer func() {
			sw.Done()
			fmt.Println("go2 exit")
		}()
		fmt.Println("go2")
		ch, ok := <-chch
		fmt.Println("go2 通过chch拿到的ch", chch, ok)
		time.Sleep(time.Second * 3)
		val := <-ch
		fmt.Println("go2 通过ch拿到的值", val)
	}()

	sw.Wait()
}
