package slice

import "testing"

func Test_nilSliceWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试nil slice"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nilSliceWork()
		})
	}
}
