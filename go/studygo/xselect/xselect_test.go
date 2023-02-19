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
