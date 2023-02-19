package xchannels

func WriteN2Chan(ch chan int, n int) {
	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch)
	//chan不需要主动关闭，会被垃圾回收
	//只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。
	//发送方如果不close，会产生all goroutines asleep - deadlock!
	//close之后再往通道中写入，会panic:send on closed channel
}

func Write2Chan(ch chan int, val int) {
	ch <- val
}
