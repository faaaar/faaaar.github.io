package leetcode

import "testing"

func TestNumDecodings091(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"0", args{"12"}, 2},
		{"1", args{"226"}, 3},
		{"2", args{"13524"}, 4},
		{"3", args{"130579246810"}, 0},
		{"4", args{"10"}, 1},
		{"5", args{"101"}, 1},
		{"6", args{"13510"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumDecodings091(tt.args.s); got != tt.want {
				t.Errorf("NumDecodings091() = %v, want %v", got, tt.want)
			}
		})
	}
}
