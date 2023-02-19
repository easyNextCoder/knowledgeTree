package slice

import "testing"

func TestStructSliceWorkWrapper(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "测试通过参数传递结构体指针中的slice，改变之后的结果"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StructSliceWorkWrapper()
		})
	}
}

func TestStructSlicePWorkWrapper(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "测试通过参数传递结构体指针中的slice的指针，改变之后的结果"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StructSlicePWorkWrapper()
		})
	}
}

func Test_changeInForRange(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试在for range中改变slice中的内容"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			changeInForRange()
		})
	}
}
