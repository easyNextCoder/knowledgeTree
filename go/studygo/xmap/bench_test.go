package xmap

import "testing"

func Test_mapInitBench(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试make map和用数组的时间差距"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapInitBench()
		})
	}
}
