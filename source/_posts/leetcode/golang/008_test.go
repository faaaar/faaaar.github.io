package leetcode

import "testing"

func TestMyAtoi008(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"0", args{"42"}, 42},
		{"1", args{"-42"}, -42},
		{"2", args{"4193 with words"}, 4193},
		{"3", args{"words and 987"}, 0},
		{"4", args{"-91283472332"}, -2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyAtoi008(tt.args.str); got != tt.want {
				t.Errorf("MyAtoi008() = %v, want %v", got, tt.want)
			}
		})
	}
}
