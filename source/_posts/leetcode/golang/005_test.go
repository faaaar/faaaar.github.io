package leetcode

import (
	"strings"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
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
			got := LongestPalindrome(tt.args.s)
			isok := false
			for _, want := range wantList {
				if got == want {
					isok = true
					break
				}
			}

			if !isok {
				t.Errorf("LongestPalindrome() = %v, want %v", got, wantList)
			}

		})
	}
}
