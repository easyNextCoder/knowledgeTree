package xlock

import "testing"

func Test_rwmutexWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "加两把读锁"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RWMutexWork()
		})
	}
}

func TestRWMutexWork1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "加两把写锁"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RWMutexWork1()
		})
	}
}

func TestRWMutexWork2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "两把写锁 抢占资源"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RWMutexWork2()
		})
	}
}

func TestRWMutexWork3(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "一把读锁一把写锁"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RWMutexWork3()
		})
	}
}

func TestRWMutexUnlockOnly(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试不加锁，只解锁时候的情况"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RWMutexUnlockOnly()
		})
	}
}

func TestRWMutexLockUnlock(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "使用defer 实现lock和unlock"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RWMutexLockUnlock()
		})
	}
}
