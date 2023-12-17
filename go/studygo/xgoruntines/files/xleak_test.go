package files

import "testing"

func Test_xleak(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试协程泄漏"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xleak()
		})
	}
}
