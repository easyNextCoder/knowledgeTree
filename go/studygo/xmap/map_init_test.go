package xmap

import "testing"

func Test_mapInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试map的初始化"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapInit()
		})
	}
}
