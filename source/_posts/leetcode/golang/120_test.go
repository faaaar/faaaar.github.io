package leetcode

import "testing"

func TestMinimumTotal120(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"0", args{[][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumTotal120(tt.args.triangle); got != tt.want {
				t.Errorf("MinimumTotal120() = %v, want %v", got, tt.want)
			}
		})
	}
}
