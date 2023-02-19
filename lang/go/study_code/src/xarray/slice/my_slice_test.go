package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewArray(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want arraySlice
	}{
		{name: "test1", args: args{n: 2}, want: NewArray(2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArray(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkOnNewArray(t *testing.T) {
	type args struct {
		t *testing.T
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test WorkOnNewArray", args: args{t: t}}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WorkOnNewArray(tt.args.t)
		})
	}
}

func Test_arraySlice_get(t *testing.T) {
	type fields struct {
		ptr []int
		len int
		cap int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{name: "test1", fields: fields(NewArray(1)), args: args{0}, want: 0}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &arraySlice{
				ptr: tt.fields.ptr,
				len: tt.fields.len,
				cap: tt.fields.cap,
			}
			if got := p.get(tt.args.n); got != tt.want {
				t.Errorf("get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arraySlice_put(t *testing.T) {
	type fields struct {
		ptr []int
		len int
		cap int
	}
	type args struct {
		n   int
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "test1", fields: fields(NewArray(2)), args: args{1, 2}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &arraySlice{
				ptr: tt.fields.ptr,
				len: tt.fields.len,
				cap: tt.fields.cap,
			}
			p.put(tt.args.n, tt.args.val)
			if got := p.get(tt.args.n); got != tt.args.val {
				t.Errorf("get() = %v, want %v", got, tt.args.val)
			}
		})
	}
}

func Test_myAppend(t *testing.T) {
	type args struct {
		as    arraySlice
		value int
	}
	wantVal := NewArray(1)
	wantVal.put(0, 1)

	tests := []struct {
		name string
		args args
		want arraySlice
	}{
		// TODO: Add test cases.
		{name: "test myAppend", args: args{as: NewArray(0), value: 1}, want: wantVal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAppend(tt.args.as, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("myAppend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myCap(t *testing.T) {
	type args struct {
		as arraySlice
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "test myCap", args: args{as: NewArray(1)}, want: 1},
		{name: "test myCap", args: args{as: NewArray(0)}, want: 0},
		{name: "test myCap", args: args{as: NewArray(100)}, want: 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myCap(tt.args.as); got != tt.want {
				t.Errorf("myCap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myLen(t *testing.T) {
	type args struct {
		as arraySlice
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{name: "test myLen", args: args{as: NewArray(1)}, want: 1},
		{name: "test myLen", args: args{as: NewArray(0)}, want: 0},
		{name: "test myLen", args: args{as: NewArray(100)}, want: 100},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myLen(tt.args.as); got != tt.want {
				t.Errorf("myLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mySlice(t *testing.T) {
	type args struct {
		as arraySlice
		l  int
		r  int
	}

	testArr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	testArr2 := testArr1[:]
	testArr3 := testArr1[0:4]
	testArr4 := testArr1[1:]

	arraySlice1 := arraySlice{ptr: testArr1, len: len(testArr1), cap: cap(testArr1)}
	arraySlice2 := arraySlice{ptr: testArr2, len: len(testArr2), cap: cap(testArr2)}
	arraySlice3 := arraySlice{ptr: testArr3, len: len(testArr3), cap: cap(testArr3)}
	arraySlice4 := arraySlice{ptr: testArr4, len: len(testArr4), cap: cap(testArr4)}

	tests := []struct {
		name string
		args args
		want arraySlice
	}{
		{name: "test mySlice1", args: args{as: arraySlice1, l: 0, r: arraySlice1.len}, want: arraySlice1},
		{name: "test myslice2", args: args{as: arraySlice1, l: 0, r: len(testArr1)}, want: arraySlice2},
		{name: "test myslice3", args: args{as: arraySlice1, l: 0, r: 4}, want: arraySlice3},
		{name: "test myslice4", args: args{as: arraySlice1, l: 1, r: len(testArr1)}, want: arraySlice4},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySlice(tt.args.as, tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testChangeArr(t *testing.T) {
	type args struct {
		arr []int
	}
	testArr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	destArr1 := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "testArr1", args: args{arr: testArr1}, want: destArr1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testChangeArr(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testChangeArr() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(testArr1, destArr1) {
				t.Errorf("testArr1 = %v, destArr1 %v", testArr1, destArr1)
			}
		})
	}
}

func Test_useSliceAfterMake(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "use slice after make"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useSliceAfterMake()
		})
	}
}

func Test_workOnMake(t *testing.T) {
	workOnMake()

}

const (
	a = iota
	b
	c
	e

	g = 99
)

func Test_slice(t *testing.T) {
	slice()
	fmt.Println(a, b, c, e, g)

	if 10 > 5 {
		x := 10
		if x*5 < 100 {
			goto addEnd
		}

	}

	fmt.Println("we break")
addEnd:
	fmt.Print("here")
}

func Test_testMax(t *testing.T) {
	testMax()

}
