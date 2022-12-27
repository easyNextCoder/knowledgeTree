package xchannels

import "fmt"

func BlockingRead(in <-chan int) int {
	fmt.Println("BlockingRead start")
	res := <-in
	fmt.Println("BlockingRead end")
	return res
}

func BlockingRead2(in <-chan int) (int, bool) {
	fmt.Println("NonBlockingRead start")
	val, ok := <-in
	if ok {
		fmt.Println("BlockingRead2 ok = true")
	}
	fmt.Println("NonBlockingRead end")
	return val, ok
}
