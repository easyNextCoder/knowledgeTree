package yservices

import (
	"fmt"
	"testing"
	"time"
)

type X struct {
	TimerEntrust <-chan time.Time
	Timer2       <-chan time.Time
}

func TestInsert(t *testing.T) {

	r := new(X)

	go func() {
		//time.Sleep(2 * time.Second)
		fmt.Println("r.TimerEntrust", r.TimerEntrust)
		fmt.Println("-a")
		r.TimerEntrust = time.After(2 * time.Second)
		r.Timer2 = time.After(3 * time.Second)
		time.AfterFunc(time.Second*3, func() {
			fmt.Println("3 seconds done!")
		})
		go func() {

			for {
				//fmt.Println("a")
				select {
				case <-r.TimerEntrust:
					fmt.Println("hello", r.TimerEntrust)
					//<-time.After(time.Second * 6)
					fmt.Println("hello done")
				case <-r.Timer2:
					fmt.Println("here")
					return
				}

			}
		}()

	}()

	time.Sleep(time.Second * 10)

}

func TestUseYorm(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试用xorm增删改查"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseYorm()
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "测试更新"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Update()
		})
	}
}

func Test_insertUser_1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试insertOne和insertMuli之后结构体的id是否会被加上"}, //最终结果是不会加上
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertMulti()
		})
	}
}

func Test_insert(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试insert之后结构体的id是否会被加上"}, //最终结果是会加上
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insert()
		})
	}
}
