package xchannels

import "testing"

func Test_do(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "ch as lock"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			do()
		})
	}
}
