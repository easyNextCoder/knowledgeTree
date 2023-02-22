package xjson

import "testing"

func Test_ginWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试对gin.H结构体变量的unmarshal"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ginWork()
		})
	}
}
