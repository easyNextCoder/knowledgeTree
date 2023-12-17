package xmap

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	sync.RWMutex
	m map[string]int
}

func Unsafe_map() {
	c := make(map[string]int)
	go func() { //开一个goroutine写map
		for j := 0; j < 1000000; j++ {
			c[fmt.Sprintf("%d", j)] = j
		}
	}()
	go func() { //开一个goroutine读map
		for j := 0; j < 1000000; j++ {
			fmt.Println(c[fmt.Sprintf("%d", j)])
		}
	}()
	time.Sleep(time.Second * 20)
}

func Safe_map() {
	var c = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	go func() { //开一个goroutine写map
		for j := 0; j < 1000000; j++ {
			c.Lock()
			c.m[fmt.Sprintf("%d", j)] = j
			c.Unlock()
		}
	}()
	go func() { //开一个goroutine读map
		for j := 0; j < 1000000; j++ {
			c.RLock()
			fmt.Println(c.m[fmt.Sprintf("%d", j)])
			c.RUnlock()
		}
	}()
	time.Sleep(time.Second * 20)
}

type SMap struct {
	m  map[int]int
	ch chan int
}

var smap SMap = SMap{
	m:  make(map[int]int),
	ch: make(chan int),
}

func (s *SMap) Read(key int) int {
	for {
		<-s.ch
		ret := s.m[key]
		s.ch <- 1
		return ret
	}

}

func (s *SMap) Write(key int, val int) {
	for {
		<-s.ch
		s.m[key] = val
		s.ch <- 1
		return
	}
}

func chan_map() {

	go func() {

		for i := 0; i < 1000; i++ {
			res := smap.Read(i)
			fmt.Println("read", i, res)
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			smap.Write(i, -i)
			fmt.Println("writ", i)
		}

	}()
	time.Sleep(time.Second * 3)
	smap.ch <- 0
	time.Sleep(time.Second * 10)
	fmt.Println("final map is ", smap.m)
	time.Sleep(time.Second * 3)
}

func Safe_map_chan() {
	var c = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	var ch chan int = make(chan int)

	go func() {
		j := 0 //开一个goroutine写map
		for {
			//fmt.Println("first", j)
			//<-ch
			c.m[fmt.Sprintf("%d", j)] = j
			//ch <- 1
			fmt.Println("first enter 1", j)
			j++
		}

		//for {
		//	//select {
		//	//case
		//	val := <-ch
		//	fmt.Println("hello", val)
		//	//}
		//	fmt.Println("default")
		//}

	}()
	go func() { //开一个goroutine读map
		j := 0
		for {

			//<-ch
			fmt.Println("map val:", c.m[fmt.Sprintf("%d", j-1)])
			//ch <- 1
			fmt.Println("first enter 2", j)
			j++
		}

	}()
	ch <- 1
	time.Sleep(2 * time.Second)
	//ch <- 2
	//for j := 0; j < 1000000; j++ {
	//	go func() { //开一个goroutine写map
	//		<-ch
	//		c.m[fmt.Sprintf("%d", j)] = j
	//		ch <- 1
	//	}()
	//}
	//
	//for j := 0; j < 1000000; j++ {
	//	go func() { //开一个goroutine读map
	//		<-ch
	//		fmt.Println(c.m[fmt.Sprintf("%d", j)])
	//		ch <- 1
	//	}()
	//}
	time.Sleep(time.Second * 5)
}

// 并发删除和写会panic
func xmapWork() {
	m := make(map[int]int)

	go func() {
		for {
			for i := 0; i < 10; i++ {
				m[i] = 0
			}
		}

	}()

	go func() {
		for {
			for i := 0; i < 10; i++ {
				delete(m, i)
			}
		}

	}()
	time.Sleep(10)
}
