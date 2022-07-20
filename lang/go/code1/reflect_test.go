package main

import "testing"

func Test_reflectNew(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "test reflect new"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reflectNew()
		})
	}
}

func Test_reflectNew11(t *testing.T) {
	type args struct {
		a *A
	}
	x := &A{
		AA: 2,
		AB: "xuyongkang",
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "test reflect new1", args: args{a: x}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reflectNew1(tt.args.a)
		})
	}
}
