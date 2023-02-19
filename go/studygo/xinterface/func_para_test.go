package xinterface

import "testing"

func Test_interfaceWorkWrapper(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试将结构体转为interface传入函数"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interfaceWorkWrapper()
		})
	}
}
