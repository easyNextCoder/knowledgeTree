package xchannels

import (
	"fmt"
	"sync"
	"time"
)

func lockAhead() { //pass 先lock没有问题
	var mu sync.Mutex

	//mu.Lock()
	go func() {

		fmt.Println("你好, 世界")
		time.Sleep(time.Second)
		mu.Unlock()
	}()

	mu.Lock() //会等待锁从而阻塞在这里等待主线程

}

func unlockAhead() { //panic 先unlock会崩溃
	var mu sync.Mutex

	//mu.Lock()
	go func() {

		fmt.Println("你好, 世界")
		mu.Unlock()
	}()

	time.Sleep(time.Second)
	mu.Lock() //会等待锁从而阻塞在这里等待主线程

}

func lock() { //pass 同步的正确做法
	var mu sync.Mutex

	mu.Lock()
	go func() {

		fmt.Println("你好, 世界")
		mu.Unlock()
	}()

	mu.Lock() //会等待锁从而阻塞在这里等待主线程

}

func channelAsLockWithCache1() { //pass,达到预期效果

	done := make(chan int, 1)
	go func() {

		time.Sleep(time.Second)
		fmt.Println("你好，世界")
		done <- 1
	}()
	fmt.Println("main")

	<-done

}

func channelAsLockWithCache() { //no panic,main 函数提前结束无法达到预期效果

	done := make(chan int, 1)
	go func() {

		time.Sleep(time.Second)
		fmt.Println("你好，世界")
		<-done
	}()
	fmt.Println("main")

	done <- 1
}

func channelAsLock() { //pass

	done := make(chan int) //如果这个channel 加缓存之后，main就不一定等go程了
	go func() {
		fmt.Println("你好，世界")
		time.Sleep(time.Second)
		<-done
	}()
	fmt.Println("main")
	done <- 1
}

func channelAsLock2() { //pass
	done := make(chan int)
	go func() {
		fmt.Println("你好，世界")
		time.Sleep(time.Second)
		done <- 1 //这样以来不会导致死锁

	}()

	fmt.Println("main")
	//done <- 1
	<-done
}

func channelAsLock2_1() { //pass
	done := make(chan int, 0)
	go func() {
		fmt.Println("你好，世界")
		time.Sleep(time.Second)
		<-done //这样以来不会导致死锁
	}()

	fmt.Println("channelAsLock2 main")
	done <- 1

}

func channelAsLock3() { //pass
	done := make(chan int, 1)

	fmt.Println("channelAsLock3 main")
	done <- 1
	<-done
}

func channelAsLock4() { //deadlock
	done := make(chan int, 0)

	fmt.Println("channelAsLock3 main")
	done <- 1
	<-done
}

func channelAsLock4_0() { //panic
	done := make(chan int, 0)

	fmt.Println("channelAsLock4_0 main")
	done <- 1
	<-done
}
func channelAsLock4_1() { //deadlock
	done := make(chan int, 0)

	fmt.Println("channelAsLock3 main")
	<-done
	done <- 1
}

func channelAsLock5() {
	done := make(chan int, 10) // 带 10 个缓存

	// 开 N 个后台打印线程
	for i := 0; i < cap(done); i++ {
		i := time.Duration(i)
		go func() {
			time.Sleep(i * time.Second / 2)
			fmt.Println(i, "你好, 世界")
			done <- 1
			fmt.Println("ingo len(done) cap(donoe)", len(done), cap(done))
		}()
		fmt.Println("for len(done) cap(donoe)", len(done), cap(done))
	}

	// 等待 N 个后台线程完成
	for i := 0; i < cap(done)-5; i++ { //这样的话就只能等待5个就结束主线程了
		<-done
	}
	fmt.Println("final len(done) cap(done)", len(done), cap(done))
}

func myChannelAsLock5() {
	done := make(chan int, 10)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			time.Sleep(time.Second * 1)
			fmt.Println("this is ", i, " printing")
			done <- i
		}()
	}

	for i := 0; i < 5; i++ {
		thei, ok := <-done
		if ok {
			fmt.Println(" the ", thei, " print done")
		}
	}

	fmt.Println("finally print done")
}

func channelUnread() {
	chun := make(chan int, 10)
	cnt := 11 //写入超过缓存数量的元素会deadlock
	for ; cnt > 0; cnt-- {
		chun <- 10
	}
	fmt.Println("用len获取通道中未读取的元素个数", len(chun))

	cnt = 100              //读出超过缓存数量的元素同样deadlock
	for ; cnt > 0; cnt-- { //panic
		res, ok := <-chun
		fmt.Println("reading channel content", ok, res)
	}
}

// 生产者: 生成 factor 整数倍的序列
func Producer(factor int, out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i * factor //往closed的通道中发送数据会panic
	}
	close(out)
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func PcWork() {
	ch := make(chan int, 64) // 成果队列

	go Producer(3, ch) // 生成 3 的倍数的序列
	go Producer(5, ch) // 生成 5 的倍数的序列
	go Consumer(ch)    // 消费生成的队列

	// 运行一定时间后退出
	time.Sleep(5 * time.Second)
}

func SelectConsumer(in <-chan int, done chan int) {
	for {
		a, ok := <-in
		time.Sleep(500 * time.Millisecond)
		if !ok {
			done <- 1
			fmt.Println("in chan is closed", ok, a)
		}
		fmt.Println("selectConsumer ok", a)
		//select {
		//case a := <-in:
		//	fmt.Println("selectProduce ok", a)
		//	//default:
		//	//	fmt.Println("channel is closed")
		//	//break
		//}
	}

}

func SelectProducer(factor int, out chan<- int) {
	for i := 0; i < 20; i++ {
		out <- i * factor
	}

	close(out)
}

func PcWorkSelect() {
	ch := make(chan int, 2) // 成果队列
	done := make(chan int, 2)

	//go SelectProducer(3, ch) // 生成 3 的倍数的序列
	go SelectProducer(5, ch)    // 生成 5 的倍数的序列
	go SelectConsumer(ch, done) // 消费生成的队列

	// 运行一定时间后退出
	<-done
}
