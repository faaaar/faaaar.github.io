package leetcode

import "testing"

func TestNumTrees096(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"0", args{3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumTrees096(tt.args.n); got != tt.want {
				t.Errorf("NumTrees096() = %v, want %v", got, tt.want)
			}
		})
	}
}
