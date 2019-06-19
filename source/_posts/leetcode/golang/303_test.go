package leetcode

import (
	"testing"
)

func TestNumArray303_SumRange303(t *testing.T) {
	type fields struct {
		Nums []int
		Sum  []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{"0", fields{[]int{-2, 0, 3, -5, 2, -1}, []int{}}, args{0, 2}, 1},
		{"0", fields{[]int{-2, 0, 3, -5, 2, -1}, []int{}}, args{2, 5}, -1},
		{"0", fields{[]int{-2, 0, 3, -5, 2, -1}, []int{}}, args{0, 5}, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor303(tt.fields.Nums)
			if got := this.SumRange303(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("NumArray303.SumRange303() = %v, want %v", got, tt.want)
			}
		})
	}
}
