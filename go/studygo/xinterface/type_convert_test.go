package xinterface

import "testing"

func Test_typeConvertWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试interface{}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typeConvertWork()
		})
	}
}
