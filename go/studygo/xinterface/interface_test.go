package xinterface

import "testing"

func Test_workonChangeInterface(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test change interface direction"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workOnChangeInterface()
		})
	}
}

func Test_testRsp(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试cmd{rsp chan<-struct}方式的命令执行"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testRsp()
		})
	}
}

func Test_testRsp2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试cmd{rsp chan<-&struct}方式的命令执行"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testRsp2()
		})
	}
}
