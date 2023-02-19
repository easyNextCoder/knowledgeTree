package xselect

import "testing"

func Test_blockWaitWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试select阻塞等待"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blockWaitWork()
		})
	}
}
