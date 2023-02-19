package slice

import "testing"

func TestSliceCopy(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "copy之前必须用make申请特定长度的数组"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SliceCopy()
		})
	}
}
