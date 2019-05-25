package leetcode

import "testing"

func TestReverse007(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"0", args{123}, 321},
		{"1", args{-123}, -321},
		{"2", args{120}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse007(tt.args.x); got != tt.want {
				t.Errorf("Reverse007() = %v, want %v", got, tt.want)
			}
		})
	}
}
