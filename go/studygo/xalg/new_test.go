package xalg

import "testing"

func Test_minimumTotalPrice(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
		price []int
		trips [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "测试图的遍历"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotalPrice(tt.args.n, tt.args.edges, tt.args.price, tt.args.trips); got != tt.want {
				t.Errorf("minimumTotalPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

//dp -1 0
//dp 0 1
//dp 1 2
//dp 2 3
//dp 2 4
func Test_minimumTotlPrice(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
		price []int
		trips [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "测试之前的结果"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotlPrice(tt.args.n, tt.args.edges, tt.args.price, tt.args.trips); got != tt.want {
				t.Errorf("minimumTotlPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
