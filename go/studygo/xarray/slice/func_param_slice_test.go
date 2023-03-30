package slice

import "testing"

func Test_workPassNilSlice(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "传递空的slice，"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workPassNilSlice()
		})
	}
}

func Test_sliceCopyWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试函数来传递slice"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceCopyWork()
		})
	}
}
