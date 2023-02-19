package xjson

import "testing"

func Test_jsonWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试json转换相关"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonWork()
		})
	}
}

func Test_nestStructMarshalWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "测试嵌套的struct"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nestStructMarshalWork()
		})
	}
}
