package xconst

import "testing"

func TestIotaWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试多个iota连用"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IotaWork()
		})
	}
}
