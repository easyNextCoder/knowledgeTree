package main

import "testing"

func Test_jsonMarshal(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test json marshal"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonMarshal()
		})
	}
}
