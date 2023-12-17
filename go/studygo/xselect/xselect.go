package xselect

import (
	"fmt"
	"time"
)

type Bus struct {
	testTimer time.Timer
	intVal    int
	arrVal    []int
	mapVal    map[int]bool
	chanVal   chan int
}

func useBus(bs *Bus) {

	bs.intVal = -1
	bs.arrVal = append(bs.arrVal, -2)
	bs.mapVal[-3] = true
	<-bs.testTimer.C
	bs.chanVal <- -4
}

func structTimer() {

	bus := new(Bus)

	go func() {
		for {
			fmt.Printf("for loop %+v", bus)
			select {

			case <-bus.testTimer.C:
				fmt.Println(bus.testTimer, bus.testTimer.C)
				fmt.Println("Timer off xxxxxxxxxxxxxxxxxxxxx")

			}
			fmt.Println("b")
		}
	}()

	yy := time.NewTimer(time.Second * 3)
	fmt.Println(yy.C)

	bus.testTimer = *yy
	//res := bus.testTimer.Stop()
	res := false
	fmt.Println(res, bus.testTimer, bus.testTimer.C)
	var s string
	fmt.Scanln(&s)
	time.Sleep(5 * time.Second)
}

func selectAndChannel() {
	ch := make(chan int)
	ch2 := make(chan int)
	//close(ch)
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("ch1")
			case <-ch2:
				fmt.Println("ch2")

			}
			fmt.Println("for run one time")
		}

		fmt.Println("go work")
	}()

	time.Sleep(5 * time.Second)
	ch2 <- 2
	time.Sleep(2 * time.Second)
	ch <- 1
	time.Sleep(5 * time.Second)
	//select {}
}

func workOnCommonChan() {
	gameStatus := make(chan int, 16) //这里设置和不设置缓存是两种不同的效果
	GameStart := 0
	RoundStart := 1
	go func() {
		for {
			select {
			case val := <-gameStatus:
				if val == GameStart {
					fmt.Println("GameStart")
					gameStatus <- RoundStart //即便是有缓存，代码没有执行完这一块仍然不会接受下一个cmd
					time.Sleep(time.Second * 2)
					fmt.Println("GameStart done!")
				} else if val == RoundStart {
					fmt.Println("RoundStart")
				}
			}
		}
	}()

	gameStatus <- GameStart

	time.Sleep(time.Second * 5)
}
