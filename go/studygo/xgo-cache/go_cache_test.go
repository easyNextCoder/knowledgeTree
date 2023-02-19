package xgo_cache

import "testing"

func Test_cacheWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试存储结构体"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cacheWork()
		})
	}
}
