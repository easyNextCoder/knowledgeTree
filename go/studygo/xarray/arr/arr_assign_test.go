package arr

import "testing"

func Test_arrAssignWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试数组的赋值"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrAssignWork()
		})
	}
}
