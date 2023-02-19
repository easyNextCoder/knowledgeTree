package slice

import "testing"

func Test_changeRangeItem(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "在item-for进行遍历的时候，item一直到只有一个，会一只复用"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			changeRangeItem()
		})
	}
}
