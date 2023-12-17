package xmap

import "testing"

func TestSafe_map(t *testing.T) {
	Safe_map()
}

func TestSafe_map_chan(t *testing.T) {
	Safe_map_chan()

}

func Test_chan_map(t *testing.T) {

	chan_map()
}

func Test_xmapWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "ceshi"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xmapWork()
		})
	}
}

func TestUnsafe_map(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试不安全的map"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Unsafe_map()
		})
	}
}
