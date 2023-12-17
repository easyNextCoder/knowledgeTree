package files

import (
	"testing"
	"time"
)

func Test_mainGo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试外层函数执行完之后，go开始执行的函数是否会丢失对象造成panic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mainGo()
			time.Sleep(time.Second * 10)
		})
	}
}
