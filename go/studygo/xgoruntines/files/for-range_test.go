package files

import "testing"

func Test_goInForRange(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试闭包函数和普通函数"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			goInForRange()
			goFuncInForRange()
		})
	}
}
