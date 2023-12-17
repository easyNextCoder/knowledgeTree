package xchannels

import "testing"

func Test_cacheReadDo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "读带缓存的空channel"}, //结果会阻塞
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nullCacheReadDo()
		})
	}
}

func Test_nullCacheWriteDo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "写入带缓存的空channel"}, //结果不会阻塞
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nullCacheWriteDo()
		})
	}
}
