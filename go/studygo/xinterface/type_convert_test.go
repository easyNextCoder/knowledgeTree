package xinterface

import "testing"

func Test_typeConvertWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试实现接口的struct的变量转换"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typeConvertWork()
		})
	}
}

func Test_typeConvertPanic(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "接口转换非自有的类型会panic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typeConvertPanic()
		})
	}
}

func Test_stringSliceConvert(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "判断interface下隐藏的类型"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringSliceConvert()
		})
	}
}
