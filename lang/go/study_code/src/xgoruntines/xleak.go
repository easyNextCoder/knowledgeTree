package xgoruntines

import (
	"fmt"
	"sync"
	"time"
)

func xleak() {

	mp := sync.Map{}
	mp.Store(1, 1)
	mp.Store("val", "ue")

	wg := sync.WaitGroup{}
	c := make(chan int, 90)

	emit := func() {
		for i := 0; i < 100; i++ {
			c <- 1
		}
	}

	time.AfterFunc(time.Second*2, emit)

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			<-c
			fmt.Println("work here", i)
		}()
	}

	wg.Wait()

}
