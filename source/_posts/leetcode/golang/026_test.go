package leetcode

import "testing"

var pairs = []struct {
	k []int
	v int
}{}

func TestRemoveDuplicates026(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{[]int{1, 1, 2}}, 2},
		{"0", args{[]int{1, 2, 3, 4, 4, 5}}, 5},
		{"0", args{[]int{5, 4, 4, 3, 2, 1}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicates026(tt.args.nums); got != tt.want {
				t.Errorf("RemoveDuplicates026() = %v, want %v", got, tt.want)
			}
		})
	}
}
