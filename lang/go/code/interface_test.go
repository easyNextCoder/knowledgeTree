package main

import "testing"

func Test_workonChangeInterface(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test change interface direction"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workOnChangeInterface()
		})
	}
}
