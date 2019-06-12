package leetcode

import "testing"

func TestIntToRoman012(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"0", args{3}, "III"},
		{"1", args{4}, "IV"},
		{"2", args{9}, "IX"},
		{"3", args{58}, "LVIII"},
		{"4", args{1994}, "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToRoman012(tt.args.num); got != tt.want {
				t.Errorf("IntToRoman012() = %v, want %v", got, tt.want)
			}
		})
	}
}
