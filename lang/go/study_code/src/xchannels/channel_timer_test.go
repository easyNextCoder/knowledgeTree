package xchannels

import (
	"fmt"
	"testing"
	"time"
)

func Test_timeAfter(t *testing.T) {
	//go invokeTimeAfter("a")
	//<-time.After(time.Second)
	//go invokeTimeAfter("b")
	//time.Sleep(5 * time.Second)
	//mp := make(map[int]map[string]interface{})
	//mp[1] = make(map[string]interface{})
	//mp[1]["xyk"] = time.Now()
	//mp[1]["name"] = 64
	//bytes, _ := json.Marshal(mp)
	//fmt.Println(string(bytes))

	var playing_timer <-chan time.Time = nil
	go func() {
		for {
			fmt.Println("first")
			select {
			case _, ok := <-playing_timer:
				if ok {
					fmt.Println("hello")
					break
				}
			}
		}
	}()
	playing_timer = time.After(1 * time.Second)
	playing_timer = time.After(2 * time.Second)
	time.Sleep(10 * time.Second)

}
