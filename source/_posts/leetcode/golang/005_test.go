package leetcode

import (
	"strings"
	"testing"
)

func TestLongestPalindrome005(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"0", args{"babad"}, "aba|bab"},
		{"1", args{"cbbd"}, "bb"},
		{"2", args{"abasbbbbbbbbbbabb"}, "bbbbbbbbbb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			wantList := strings.Split(tt.want, "|")
			got := LongestPalindrome005(tt.args.s)
			isok := false
			for _, want := range wantList {
				if got == want {
					isok = true
					break
				}
			}

			if !isok {
				t.Errorf("LongestPalindrome005() = %v, want %v", got, wantList)
			}

		})
	}
}
