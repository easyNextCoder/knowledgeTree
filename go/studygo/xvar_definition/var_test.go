package xvar_definition

import "testing"

func TestWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "测试用var直接定义数组"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Work()
		})
	}
}
