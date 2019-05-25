package leetcode

import (
	"reflect"
	"testing"
)

var _002TestData = []struct {
	L1  *ListNode
	L2  *ListNode
	Out *ListNode
}{}

func TestAddTwoNumbers002(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"0", args{NewListNode([]int{2, 4, 3}), NewListNode([]int{5, 6, 4})}, NewListNode([]int{7, 0, 8})},
		{"1", args{NewListNode([]int{1, 2, 3}), NewListNode([]int{6, 7, 8, 9})}, NewListNode([]int{7, 9, 1, 0, 1})},
		{"2", args{NewListNode([]int{5}), NewListNode([]int{5})}, NewListNode([]int{0, 1})},
		{"3", args{NewListNode([]int{0}), NewListNode([]int{7, 3})}, NewListNode([]int{7, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers002(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers002() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddTwoNumbersRecursive002(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"0", args{NewListNode([]int{2, 4, 3}), NewListNode([]int{5, 6, 4})}, NewListNode([]int{7, 0, 8})},
		{"1", args{NewListNode([]int{1, 2, 3}), NewListNode([]int{6, 7, 8, 9})}, NewListNode([]int{7, 9, 1, 0, 1})},
		{"2", args{NewListNode([]int{5}), NewListNode([]int{5})}, NewListNode([]int{0, 1})},
		{"3", args{NewListNode([]int{0}), NewListNode([]int{7, 3})}, NewListNode([]int{7, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbersRecursive002(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbersRecursive002() = %v, want %v", got, tt.want)
			}
		})
	}
}
