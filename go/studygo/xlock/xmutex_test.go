package xlock

import "testing"

func Test_mutexWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "互斥锁测试"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutexWork()
		})
	}
}

func Test_mutex2Work(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "互斥锁定过程中，崩溃不释放锁"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutexWork3()
		})
	}
}

func Test_mutexWork2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "两把互斥锁 进行同步"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutexWork2()
		})
	}
}
