package leetcode

import "testing"

func TestMaxArea011(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"0", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxArea011(tt.args.height); got != tt.want {
				t.Errorf("MaxArea011() = %v, want %v", got, tt.want)
			}
		})
	}
}
