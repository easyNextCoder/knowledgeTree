package xlock

import (
	"fmt"
	"sync"
	"time"
)

type MutexSource struct {
	mutex sync.Mutex //互斥锁 保证一段代码互斥执行 https://tour.go-zh.org/concurrency/9
	val   [10]int
}

//互斥锁连续锁定两次 panic
func mutexWork() {
	var mutex sync.Mutex

	mutex.Lock()
	mutex.Lock()

}

var sa MutexSource
var mp map[int]int = make(map[int]int)
var startVal int = 0

func write(i int) {
	sa.mutex.Lock() //每个锁定都会阻塞后来者，所以必须要进行解锁
	startVal += i
	sa.mutex.Unlock()
}

func mutexWork2() {

	sw := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		sw.Add(1)
		go func() {
			defer sw.Done()
			//fmt.Println("go1")
			write(1)
		}()
	}
	for i := 0; i < 100; i++ {
		sw.Add(1)
		go func() {
			defer sw.Done()
			//fmt.Println("go2")
			write(-1)
		}()
	}
	sw.Wait()
	fmt.Println(startVal)
}

//go1获得了互斥锁之后崩溃了并没有释放互斥锁，这个时候go2再尝试去获得互斥锁会 panic
func mutexWork3() {
	var ms MutexSource

	sw := sync.WaitGroup{}

	sw.Add(1)
	go func() {
		defer sw.Done()
		fmt.Println("go2 尝试锁定")
		ms.mutex.Lock()
		fmt.Println("go2 锁定成功")
		ms.mutex.Unlock()
	}()

	sw.Add(1)
	go func() {
		defer func() {
			defer sw.Done()
			if e := recover(); e != nil {
				fmt.Println("go1 recover", e)
			}
			time.Sleep(3 * time.Second)
			ms.mutex.Unlock() //释放锁之后，go2可以获得即可成功
		}()
		fmt.Println("go1 尝试锁定")
		ms.mutex.Lock()
		fmt.Println("go1 锁定成功")
		panic("go1 崩溃")
		ms.mutex.Unlock()

	}()

	sw.Wait()
}
