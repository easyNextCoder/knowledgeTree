package xvar_definition

import (
	"fmt"
	"sync"
	"time"
)

func varDefine() {

	//var x []int //正确可以append，之后可以读写
	//var x map[int]int  //错误读不panic， 写会发生panic
	//var x chan int //错误 发生panic
	//var x struct{}{} //正确
	//
	//y := []int{}//正确
	//y := map[int]int{}//正确
	//
	//z := make([]int, 0)//正确
	//z := make(map[int]int)//正确
	//z := make(chan int)//正确
	//

	var x <-chan time.Time

	sw := sync.WaitGroup{}
	sw.Add(1)
	go func() {
		defer sw.Done()
		fmt.Println("go1 sleep 1s")

		fmt.Println("go1 sleep 1s done")
		x = time.After(time.Second * 3)
		fmt.Println("go1 sleep 3s done")
		time.Sleep(time.Second * 4)
		fmt.Println("go1 sleep 4s done")

	}()
	sw.Add(1)
	go func() {
		defer sw.Done()
		fmt.Println("go2 enter")
		time.Sleep(time.Second * 1)
		for {
			select {
			case y := <-x:
				fmt.Println("this is y:", y)

			case <-time.After(time.Second * 1):
				fmt.Println("tick 1s")
			}
			fmt.Println("go2 done")
		}

	}()
	fmt.Println("all go done")
	sw.Wait()
}

func structDefine() {
	type TickS struct {
		timer <-chan time.Time
	}

	tick := new(TickS)

	fmt.Println(tick.timer)

	sw := sync.WaitGroup{}

	sw.Add(1)
	go func() {
		defer sw.Done()
		time.Sleep(5 * time.Second)
		tick.timer = time.After(time.Second * 2)
	}()

	sw.Add(1)
	go func() {
		defer sw.Done()
		for {
			fmt.Println(tick.timer)
			select {
			case y := <-tick.timer:
				fmt.Println("y is", y)
			case <-time.After(1 * time.Second): //定时触发
				fmt.Println("tick 1s")
			}

		}
	}()

	sw.Wait()
}

func intChan() {
	x := make(chan int)

	//var x chan int

	sw := sync.WaitGroup{}

	sw.Add(1)
	go func() {
		defer sw.Done()
		time.Sleep(time.Second * 1)
		x <- 10
	}()

	sw.Add(1)
	go func() {
		defer sw.Done()
		for {
			fmt.Println("enter for")
			select {
			case <-x:
				fmt.Println("x enter")
				return //没有return 会发生panic
			}

		}
	}()

	sw.Wait()
}
