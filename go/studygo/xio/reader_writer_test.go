package xio

import "testing"

func Test_readerWork(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "测试reader接口"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := readerWork(); (err != nil) != tt.wantErr {
				t.Errorf("readerWork() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_buildGZipReaderWork(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "测试gzip的reader接口"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := buildGZipReaderWork(); (err != nil) != tt.wantErr {
				t.Errorf("buildGZipReaderWork() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_selfReaderWork(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "测试自己写的reader"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selfReaderWork()
		})
	}
}
