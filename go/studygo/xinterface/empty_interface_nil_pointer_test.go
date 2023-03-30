package xinterface

import "testing"

func Test_emptyInterfaceNilPointer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "测试空接口和空指针是否是相等的"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emptyInterfaceNilPointer()
		})
	}
}
