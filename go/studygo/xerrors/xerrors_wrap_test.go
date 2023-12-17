package xerrors

import "testing"

func TestWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试errors"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Work()
		})
	}
}

func Test_workWrapErr(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "测wrapper ERR"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := workWrapErr(); (err != nil) != tt.wantErr {
				t.Errorf("workWrapErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
