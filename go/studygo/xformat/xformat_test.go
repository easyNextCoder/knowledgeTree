package xformat

import "testing"

func Test_precisionControl(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试精度打印"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			precisionControl()
		})
	}
}
