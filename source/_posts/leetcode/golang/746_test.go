package leetcode

import "testing"

func TestMinCostClimbingStairs746(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"0", args{[]int{10, 15, 20}}, 15},
		{"1", args{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCostClimbingStairs746(tt.args.cost); got != tt.want {
				t.Errorf("MinCostClimbingStairs746() = %v, want %v", got, tt.want)
			}
		})
	}
}
