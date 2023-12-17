package xalg

import (
	"fmt"
	"testing"
)

func Test_heapSort20231130(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "测试堆排序", args: args{arr: GenRandomArr(300, 0, 500)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heapSort20231130(tt.args.arr)
			fmt.Println("函数内数组无法赋值", tt.args.arr)
		})
	}
}
