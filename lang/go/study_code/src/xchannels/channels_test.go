package xchannels

import (
	"fmt"
	"testing"
	"time"
)

func TestReadWriteInitializedChan(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test read write initializalized chan"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadWriteInitializedChan()
		})
	}
}

func TestReadWriteUninitializedChan(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "Test read write uninitialized chan"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadWriteUninitializedChan()
		})
	}
}

func TestReadWriteInitializedChanWithBuffer(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test read write initialized chan with buffer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadWriteInitializedChanWithBuffer()
		})
	}
}

func TestChannelRead(t *testing.T) {

	out := make(chan int)
	out2 := make(chan int)
	//只写的通
	go func() {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("2 second time out")
			out <- 1
			out2 <- 2
			//default:
			//	fmt.Println("no message received")
		}

	}()

	go func() {
		BlockingRead(out)
	}()

	go func() {
		BlockingRead2(out2)
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("this is TestWork.")
}

func Test_chanSendStruct(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "测试用chan来发送结构体"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chanSendStruct()
		})
	}
}
