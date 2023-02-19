package xinterface

import "testing"

func Test_mainWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "用interface实现多态"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mainWork()
		})
	}
}
