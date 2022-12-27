package xreflect

import "testing"

func Test_xreflect(t *testing.T) {
	xreflect()

}

func Test_xreflect1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "work1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xreflect()
		})
	}
}
