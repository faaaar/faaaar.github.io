package leetcode

import "testing"

func TestUniquePathsWithObstacles063(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"0", args{[][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}}, 2},
		{"1", args{[][]int{[]int{0, 0, 0, 0}, []int{0, 1, 0, 0}, []int{0, 0, 0, 0}, []int{0, 0, 0, 0}}}, 8},
		{"2", args{[][]int{[]int{0, 1, 0, 0}, []int{0, 0, 0, 0}, []int{0, 0, 0, 0}, []int{0, 0, 0, 0}}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePathsWithObstacles063(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("UniquePathsWithObstacles063() = %v, want %v", got, tt.want)
			}
		})
	}
}
