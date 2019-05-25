package leetcode

import "testing"

func TestLengthOfLongestSubstring003(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"0", args{""}, 0},
		{"1", args{"pwwkew"}, 3},
		{"2", args{"aaaaaaa"}, 1},
		{"3", args{"dvdf"}, 3},
		{"4", args{"abba"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring003(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring003() = %v, want %v", got, tt.want)
			}
		})
	}
}
