package yservices

import "testing"

func Test_deleteAll(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "delete时候不加where条件无法删除所有的，加上where 1=1 条件之后可以无报错删除"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := deleteAll(); (err != nil) != tt.wantErr {
				t.Errorf("deleteAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
