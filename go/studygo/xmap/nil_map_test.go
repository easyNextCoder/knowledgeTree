package xmap

import "testing"

func Test_nilMapWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试用var申请nil map"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nilMapWork()
		})
	}
}

func Test_assMapWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "用:=来生成map，并测试读写"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assMapWork()
		})
	}
}
