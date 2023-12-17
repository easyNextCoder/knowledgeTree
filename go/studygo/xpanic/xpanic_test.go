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
			//panicRecover()
		})
	}
}

func Test_panicRecoverWrapper(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试defer处理panic之后，程序仍可以正常运行"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			panicRecoverWrapper()
		})
	}
}
