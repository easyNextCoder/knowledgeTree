package xbit

import "testing"

func Test_runCheck(t *testing.T) {
	runCheck()
}

func TestSetInt64(t *testing.T) {
	type args struct {
		in     uint64
		v      int
		border [2]int
	}
	tests := []struct {
		name    string
		args    args
		wantOut uint64
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "测试设置一个数中的某几位", args: args{
			in:     0,
			v:      7,
			border: [2]int{1, 4},
		}, wantOut: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := SetInt64(tt.args.in, tt.args.v, tt.args.border)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut != tt.wantOut {
				t.Errorf("SetInt64() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
