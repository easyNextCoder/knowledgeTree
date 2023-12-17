package xchannels

import "testing"

func Test_chanChan(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "通过channel来传递channels"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chanChan()
		})
	}
}
