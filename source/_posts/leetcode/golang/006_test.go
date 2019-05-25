package leetcode

import (
	"testing"
)

func TestConvert006(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"0", args{"LEETCODEISHIRING", 1}, "LEETCODEISHIRING"},
		{"1", args{"LEETCODEISHIRING", 3}, "LCIRETOESIIGEDHN"},
		{"2", args{"LEETCODEISHIRING", 4}, "LDREOEIIECIHNTSG"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert006(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("Convert006() = %v, want %v", got, tt.want)
			}
		})
	}
}
