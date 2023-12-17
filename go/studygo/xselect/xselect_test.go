package xselect

import "testing"

func Test_selectAndChannel(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试用select和不用select，channel对go程的阻塞效果"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selectAndChannel()
		})
	}
}

func Test_structTimer(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试结构体中的timer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			structTimer()
		})
	}
}

func Test_workOnCommonChan(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "执行一个chan case的时候，向这个chan发送相同的case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workOnCommonChan()
		})
	}
}
