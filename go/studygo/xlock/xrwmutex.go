package xlock

import (
	"fmt"
	"sync"
	"time"
)

type RWMutexSource struct {
	rw  sync.RWMutex //读写锁 保证并发读写安全，已经被锁定的情况下新来的将被阻塞
	val [10]int
}

func RWMutexWork() {
	var sa RWMutexSource
	sa.rw.RLock()
	sa.rw.RLock()
}

func RWMutexWork1() {
	var sa RWMutexSource
	sa.rw.Lock()
	sa.rw.Lock()
}

var slock sync.RWMutex

func RWMutexUnlockOnly() {

	var a, b int
	a = b + 1
	if a > b {
		fmt.Println("不会跑defer 语句")
		return
	}
	defer func() {
		fmt.Println("这里会调用defer的unlock")
		slock.Unlock()
	}()
}

func RWMutexLockUnlock() {
	defer slock.Unlock()
	slock.Lock()
}

func RWMutexWork2() {
	//测试写锁

	var sa RWMutexSource
	mp := map[int]int{}
	sw := sync.WaitGroup{}

	for i := 0; i < 10000; i++ {
		sw.Add(1)
		go func() {
			defer sw.Done()
			sa.rw.Lock()
			mp[i] = i
			sa.rw.Unlock()
		}()
	}
	for i := 0; i < 10000; i++ {
		sw.Add(1)
		go func() {
			defer sw.Done()
			sa.rw.Lock()
			mp[i] = -i
			sa.rw.Unlock()
		}()
	}
	sw.Wait()
}
func RWMutexWork3() {
	//测试写锁

	var sa RWMutexSource
	mp := map[int]int{}
	sw := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		sw.Add(1)
		go func() {
			defer sw.Done()
			sa.rw.RLock()
			fmt.Println(mp[i])
			sa.rw.RUnlock()
		}()
	}
	for i := 0; i < 10000; i++ {
		sw.Add(1)
		go func() {
			defer sw.Done()
			sa.rw.Lock()
			mp[i] = -i
			sa.rw.Unlock()
		}()
	}
	sw.Wait()
}

func WaitLock() {
	var l sync.RWMutex
	sg := sync.WaitGroup{}
	sg.Add(1)
	time.AfterFunc(time.Second*1, func() {

		go func() {
			defer l.Unlock()
			l.Lock()
			fmt.Println("later")
			sg.Done()
		}()
	})

	sg.Add(1)
	go func() {
		defer l.Unlock()
		l.Lock()
		fmt.Println("first sleep")
		time.Sleep(3 * time.Second)
		fmt.Println("first sleep done")
		sg.Done()
	}()

	sg.Wait()
	fmt.Println("final done!")
}
