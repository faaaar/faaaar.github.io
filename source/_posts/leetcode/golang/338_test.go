package leetcode

import (
	"reflect"
	"testing"
)

func TestCountBits338(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"0", args{2}, []int{0, 1, 1}},
		{"1", args{5}, []int{0, 1, 1, 2, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBits338(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountBits338() = %v, want %v", got, tt.want)
			}
		})
	}
}
