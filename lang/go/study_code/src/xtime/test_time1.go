package xtime

import (
	"fmt"
	"time"
)

func timeFunc() {
	fmt.Println(time.Date(2022, 8, 5, 0, 0, 0, 0, time.Local))
}

func timeAfterFunc() {
	f := func() {
		fmt.Println("afterFunc work..")
	}
	time.AfterFunc(time.Second*3, f)

	fmt.Println("done")
}

func play(s string) {
	fmt.Println("playing...")
	<-time.After(time.Second)
	fmt.Println(s)
}

func outPlay(s string) {
	play(s)
}

func runTimeFunc() {
	//timeFunc()
	timeAfterFunc()
	go outPlay("hello")
	<-time.After(time.Second)
	go outPlay("work")
	var a int = 1
	fmt.Printf("%x", a)
	time.Sleep(time.Second * 5)

	startTime := time.Now()
	<-time.After(2 * time.Second)
	endTime := time.Now()
	var val int64 = 199
	var mp map[int]int
	mp[0] = 0
	fmt.Printf("%v, %d, %d", endTime.Sub(startTime).Seconds(), val, len(mp))

}

func testTimeAfter() {
	fmt.Println("we first print")
	<-time.After(time.Second * 5)
	fmt.Println("we print")
}
