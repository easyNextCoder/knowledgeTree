package arr

import "testing"

func Test_arrCopyWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "测试数组的拷贝"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrCopyWork()
		})
	}
}
