package arr

import "testing"

func Test_arrAllocWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试var定义全局数组"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrAllocWork()
		})
	}
}

func Test_work(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试图的遍历"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			work()
		})
	}
}
