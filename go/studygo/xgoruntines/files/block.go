package files

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func RunAfterBlock() {
	go func() {
		v, ok := <-ch
		fmt.Println("receive done", v, ok)

	}()

	time.Sleep(time.Second * 2)

	go func() {
		ch <- 1
		fmt.Println("send done")
	}()
	time.Sleep(time.Second * 2)
}
