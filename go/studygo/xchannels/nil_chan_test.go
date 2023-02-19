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
