package xalg

import "testing"

func Test_sortSliceWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试使用sort逆向排序"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortSliceWork()
		})
	}
}
