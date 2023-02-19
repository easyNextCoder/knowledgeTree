package xtime

import "testing"

func Test_runTimeFunc(t *testing.T) {

	runTimeFunc()

}

func Test_testTimeAfter(t *testing.T) {
	testTimeAfter()
}

func Test_timeConvert(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试时间转换之后打印的格式"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeConvert()
		})
	}
}
