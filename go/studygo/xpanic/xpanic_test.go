package xpanic

import "testing"

func Test_directPanic(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试panic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			directPanic()
		})
	}
}

func Test_panicRecover(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试panicRecover"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			panicRecover()
		})
	}
}
