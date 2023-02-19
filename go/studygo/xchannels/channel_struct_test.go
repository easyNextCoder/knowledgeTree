package xchannels

import "testing"

func Test_channelStructWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试使用channel传递结构体"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			channelStructWork()
		})
	}
}

func Test_structPointer(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试用channel传递结构体指针"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			structPointer()
		})
	}
}
