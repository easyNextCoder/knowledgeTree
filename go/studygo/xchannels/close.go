package xchannels

import "fmt"

func closeCh() {
	done := make(chan int)
	msg := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			msg <- i + 100
		}
		close(msg) //没有这个close就会产生deadlock
	}()

	go func() {
		for v := range msg {
			fmt.Println(v)
		}

		v1 := <-msg
		v2 := <-msg
		v3 := <-msg
		//多次读写已关闭通道并不会阻塞

		fmt.Println("read closed chan:", v1, v2, v3)

		done <- 1

	}()

	<-done
}
