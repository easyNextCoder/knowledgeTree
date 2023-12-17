package xreflect

import "testing"

func Test_buildin_func(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试TypeOf函数"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TypeOfFunc()
		})
	}
}

func TestValueOfFunc(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试ValueOf函数"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ValueOfFunc()
		})
	}
}
