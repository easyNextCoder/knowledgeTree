package xvar_definition

import "testing"

func Test_newMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "new map"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newMap()
		})
	}
}

func Test_newChan(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "new chan"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newChan()
		})
	}
}

func Test_newStruct(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "new struct"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newStruct()
		})
	}
}

func Test_newArr(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "new arr"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newArr()
		})
	}
}
