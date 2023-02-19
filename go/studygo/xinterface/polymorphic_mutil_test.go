package xinterface

import "testing"

func Test_playGround(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试struct实现多接口"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			playGround()
		})
	}
}
