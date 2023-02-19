package slice

import "testing"

func Test_slicePointerWork2(t *testing.T) {
	type args struct {
		ps *[]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slicePointerWork2(tt.args.ps)
		})
	}
}
