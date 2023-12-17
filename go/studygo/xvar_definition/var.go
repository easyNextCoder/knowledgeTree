package xvar_definition

import (
	"fmt"
	"time"
)

func varArr() {

	var intArr []int
	intArr = append(intArr, 1) //ok

}

func varArrPtr() {

	var intPtrArr *[]int
	*intPtrArr = append(*intPtrArr, 1) //panic

}

func varMap() {

	var mpr map[int]bool = map[int]bool{} //ok
	mpr[-1] = true

	var mp map[int]bool //panic
	mp[-1] = true
}

func varChan() {
	var ch chan int
	ch <- 1 //panic
}

type Bus struct {
	testTimer time.Timer
	intVal    int
	arrVal    []int
	mapVal    map[int]bool
	chanVal   chan int
}

func varStruct() {
	var bs Bus

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic1")
			}
		}()
		bs.intVal = -1
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic2")
			}
		}()
		bs.arrVal = append(bs.arrVal, -2)
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic3")
			}
		}()
		bs.mapVal[-3] = true //panic
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic4")
			}
		}()
		fmt.Println("bs.testTimer.C:", bs.testTimer.C)
		<-bs.testTimer.C //nil chan 会一直阻塞，在go程中不会Panic，如果不是则会产生死锁，
		fmt.Println("execute done!")
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic5")
			}
		}()
		fmt.Println("bs.chanVal:", bs.chanVal)
		bs.chanVal <- -4 ////nil chan 一直是阻塞的无法发送也无法接受，会产生死锁但是不会panic
		fmt.Println("execute done!")
	}()

	time.Sleep(time.Second * 1)

}
