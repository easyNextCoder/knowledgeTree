package xchannels

import "testing"

func Test_nilChannelWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "测试用var来申请chan"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nilChannelWork()
		})
	}
}

func Test_notNilChannelWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "声明nil chan -> 赋值 -> 使用"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notNilChannelWork()
		})
	}
}
