package xvar_definition

import "testing"

func Test_varArr(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "用var申请数组"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			varArr()
		})
	}
}

func Test_varArrPtr(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "用var申请数组指针"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			varArrPtr()
		})
	}
}

func Test_varMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "用var申请map"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			varMap()
		})
	}
}

func Test_varChan(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "用var申请chan"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			varChan()
		})
	}
}

func Test_varStruct(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "用var申请结构体"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			varStruct()
		})
	}
}
