package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum001(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"0", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"1", args{[]int{2, 7, 11, 15, 17}, 19}, []int{0, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum001(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum001() = %v, want %v", got, tt.want)
			}
		})
	}
}
