package xredis

import "testing"

func Test_redisWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试将结构体json到redis中，并转换回来"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisWork()
		})
	}
}

func Test_redisType(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试redis类型"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisType()
		})
	}
}

func Test_bigKey(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试大key"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bigKey()
		})
	}
}

func Test_runWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试获得bigKey"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runWork()
		})
	}
}

func Test_stringOperate(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试redis字符串操作"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringOperate()
		})
	}
}

func Test_listOperate(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试redis列表操作"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listOperate()
		})
	}
}

func Test_hashOperate(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试redis哈希表操作"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashOperate()
		})
	}
}

func Test_setOperate(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试redis集合操作"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setOperate()
		})
	}
}

func Test_sortedSetOperate(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试redis有序集合操作"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortedSetOperate()
		})
	}
}
