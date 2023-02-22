package yservices

import (
	"fmt"
	"testing"
	"time"
)

type X struct {
	TimerEntrust <-chan time.Time
}

func TestInsert(t *testing.T) {

	r := new(X)

	go func() {
		//time.Sleep(2 * time.Second)
		go func() {

			for {
				fmt.Println("a")
				select {
				case <-r.TimerEntrust:
					fmt.Println("hello")
					return
				}
				fmt.Println("b")
			}
		}()
		fmt.Println("-a")
		r.TimerEntrust = time.After(2 * time.Second)

		fmt.Println("xxx")
		r.TimerEntrust = time.After(5 * time.Second)
		fmt.Println("end")
	}()

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
