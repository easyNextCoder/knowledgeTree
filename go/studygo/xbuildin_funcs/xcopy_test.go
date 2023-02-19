package xbuildin_funcs

import "testing"

func Test_copyStructLen0Arr(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "new struct的时候slice会被初始化；copy可以操作长度为0的slice"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			copyStructLen0Arr()
		})
	}
}
