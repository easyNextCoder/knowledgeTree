package arr

import "testing"

func Test_arrFuncWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "函数数组的初始化数值是否是nil?"}, //是的
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrFuncWork()
		})
	}
}
