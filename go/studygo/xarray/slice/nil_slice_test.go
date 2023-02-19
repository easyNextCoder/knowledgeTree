package slice

import "testing"

func Test_resetVarSliceToNil(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试是否可以向append(nil, x)"}, //可以append nil
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetVarSliceToNil()
		})
	}
}
