package xerrors

import "testing"

func Test_defineErrIsTimeout(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "判断一个err是否是timeout的err"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defineErrIsTimeout()
		})
	}
}
