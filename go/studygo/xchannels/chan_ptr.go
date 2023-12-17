package xchannels

import (
	"fmt"
	"time"
)

type RobotsRunner struct {
	RobotTimer1      <-chan time.Time
	RobotTimer2      <-chan time.Time
	RobotTimer3      <-chan time.Time
	RobotTimer4      <-chan time.Time
	RobotActiveLevel []int
	Uid2Timer        map[int64]*<-chan time.Time
	signChan         chan int
}

func RobotsWork() {
	r := new(RobotsRunner)
	r.Uid2Timer = map[int64]*<-chan time.Time{}
	r.Uid2Timer[1] = &r.RobotTimer1
	r.Uid2Timer[2] = &r.RobotTimer2
	r.Uid2Timer[3] = &r.RobotTimer3
	r.Uid2Timer[4] = &r.RobotTimer4
	r.signChan = make(chan int, 32)
	go r.RobotsGo()
	r.signChan <- 1
	time.Sleep(time.Second * 1)
	r.signChan <- 2
	time.Sleep(time.Second * 1)
	r.signChan <- 3
	time.Sleep(time.Second * 1)
	r.signChan <- 1
	time.Sleep(time.Second * 1)

	time.Sleep(time.Second * 10)
}

func (r *RobotsRunner) RobotsGo() {
	for {
		select {
		case v := <-r.signChan:
			fmt.Println(v, "signchan work")
			*r.Uid2Timer[int64(v)] = time.After(time.Second)
		case v := <-r.RobotTimer1:
			fmt.Println("timer 1 working", v)
		case v := <-r.RobotTimer2:
			fmt.Println("timer 2 working", v)
		case v := <-r.RobotTimer3:
			fmt.Println("timer 3 working", v)
		}
	}

	time.AfterFunc(time.Second*3, func() {
		fmt.Println("-1s")
	})
	time.AfterFunc(time.Second*2, func() {
		fmt.Println("2s")
	})
	time.Sleep(time.Second * 5)
}
