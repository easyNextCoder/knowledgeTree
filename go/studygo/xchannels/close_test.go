package xchannels

import "testing"

func Test_closeCh(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试close函数的作用"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			closeCh()
		})
	}
}
