package xchannels

import "testing"

func TestReadWrite(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "无缓存管道读写"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadWrite()
		})
	}
}
