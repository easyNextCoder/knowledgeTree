package xstruct

import "testing"

func Test_testVar(t *testing.T) {
	testVar()

}

func Test_varFuncReceivePointer(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "变量接受器的函数，是否能够接受指针的调用"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			varFuncReceivePointer()
		})
	}
}

func Test_useVarAsReceiver(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试用过变量作为函数的接受体，当改变结构体内部值之后原变量值是否改变"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useVarAsReceiver()
		})
	}
}

func Test_usePointerAsReceiver(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试用指针作为函数的接受体，当改变结构体内部值之后原变量值是否改变"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usePointerAsReceiver()
		})
	}
}
