package files

import "testing"

func Test_closuresWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "time.AfterFunc 闭包函数"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			closuresWork()
		})
	}
}

func Test_outerInner(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试闭包的函数代入"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outerInner()
		})
	}
}
