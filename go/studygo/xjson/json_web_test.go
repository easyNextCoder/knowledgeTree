package xjson

import "testing"

func Test_jsonMarshalWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试jsonMarshal"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonMarshalWork()
		})
	}
}
