package leetcode

import "testing"

func TestIsPalindrome009(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"0", args{121}, true},
		{"1", args{-121}, false},
		{"2", args{10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome009(tt.args.x); got != tt.want {
				t.Errorf("IsPalindrome009() = %v, want %v", got, tt.want)
			}
		})
	}
}
