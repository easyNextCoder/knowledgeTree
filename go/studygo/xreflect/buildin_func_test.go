package xreflect

import "testing"

func Test_buildin_func(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "反射测试1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildin_func()
		})
	}
}
