package leetcode

import "testing"

func TestMinPathSum064(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"0", args{[][]int{[]int{1, 3, 1}, []int{1, 5, 1}, []int{4, 2, 1}}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinPathSum064(tt.args.grid); got != tt.want {
				t.Errorf("MinPathSum064() = %v, want %v", got, tt.want)
			}
		})
	}
}
