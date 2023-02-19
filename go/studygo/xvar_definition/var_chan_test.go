package xvar_definition

import "testing"

func Test_varDefine(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试用var 定义chan"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			varDefine()
		})
	}
}

func Test_structDefine(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试用结构体申请chan"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			structDefine()
		})
	}
}

func Test_intChan(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "intChan 变量申请"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intChan()
		})
	}
}
