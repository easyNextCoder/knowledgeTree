package xgoruntines

import "testing"

func TestGoPanicThenRecover(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test go panic then recover"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GoPanicThenRecover()
		})
	}
}
