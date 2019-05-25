package leetcode

import "testing"

func TestFindMedianSortedArrays004(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"0", args{[]int{1, 2, 3}, []int{4, 5, 6}}, 3.5},
		{"1", args{[]int{1, 2}, []int{3, 4, 5}}, 3.0},
		{"2", args{[]int{1, 2}, []int{3, 4, 5, 6}}, 3.5},
		{"3", args{[]int{1, 2}, []int{3, 4, 5, 6}}, 3.5},
		{"4", args{[]int{2, 4}, []int{1, 3, 5}}, 3.0},
		{"5", args{[]int{9, 10}, []int{2, 4, 6}}, 6.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMedianSortedArrays004(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("FindMedianSortedArrays004() = %v, want %v", got, tt.want)
			}
		})
	}
}
