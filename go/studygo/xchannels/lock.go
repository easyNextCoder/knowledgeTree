package xchannels

import "fmt"

//两次上锁panic
func cLock(ch chan int) {
	ch <- 1
}

//未上锁先解锁panic
func cUnlock(ch chan int) {
	<-ch
}

func do() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		fmt.Println("outer i", i)
		go func(input int) { //这里是闭包
			for {
				cLock(ch)
				fmt.Println(input)
				cUnlock(ch)
			}
		}(i)
	}
	<-make(chan int)
}
