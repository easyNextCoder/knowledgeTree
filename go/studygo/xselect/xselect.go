package xselect

import (
	"fmt"
	"testing"
	"time"
)

type XX struct {
	testTimer time.Timer
}

func Test_alloc(t *testing.T) {
	//a := []int{1, 3, 4}
	//fmt.Println("this is a", a)
	//a = a[:0]
	//fmt.Println("this is a2", a)
	//alloc()
	xx := new(XX)
	go func() {
		for {
			fmt.Println("a")
			select {

			case <-xx.testTimer.C:
				fmt.Println(xx.testTimer, xx.testTimer.C)
				fmt.Println("Timer off xxxxxxxxxxxxxxxxxxxxx")

			}
			fmt.Println("b")
		}
	}()

	yy := time.NewTimer(time.Second * 3)
	fmt.Println(yy.C)

	xx.testTimer = *yy
	//res := xx.testTimer.Stop()
	res := false
	fmt.Println(res, xx.testTimer, xx.testTimer.C)
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
