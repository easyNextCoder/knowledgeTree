package slice

import "testing"

func TestLiteralCopyAppendWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试[]拷贝 用字面量申请"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LiteralCopyAppendWork()
		})
	}
}

func TestMakeCopyAppendWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试[]拷贝 用make申请"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeCopyAppendWork()
		})
	}
}

func Test_copyAssignment(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试copy之后赋值是否会改变原切片"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			copyAssignment()
		})
	}
}

func Test_makeAndLiteralAlloc(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试用make和字面申请切片的区别"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			makeAndLiteralAlloc()
		})
	}
}

func Test_sliceCut(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试[]不从起始位置取值，之后的内部指针"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceCut()
		})
	}
}

func Test_appendInRange(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试在range的时候向数组添加内容"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appendInRange()
		})
	}
}
