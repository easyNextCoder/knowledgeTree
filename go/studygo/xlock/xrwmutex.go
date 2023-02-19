package xlock

import (
	"fmt"
	"sync"
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
