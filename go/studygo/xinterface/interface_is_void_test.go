package xinterface

import "testing"

func Test_assignmentInterfaceWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "interface{}实际上就是一个胖指针(void)类型"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignmentInterfaceWork()
		})
	}
}
