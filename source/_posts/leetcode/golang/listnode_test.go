package leetcode

import (
	"reflect"
	"testing"
)

func TestNewListNode(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"0", args{[]int{1, 2, 3}}, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListNode(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_Compare(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
	}
	type args struct {
		l2 *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"0", fields{1, &ListNode{2, nil}}, args{&ListNode{1, &ListNode{2, nil}}}, true},
		{"1", fields{2, &ListNode{1, nil}}, args{&ListNode{1, &ListNode{2, nil}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}

			if got := l1.Compare(tt.args.l2); got != tt.want {
				t.Errorf("ListNode.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_String(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"0", fields{1, &ListNode{2, &ListNode{3, nil}}}, "1,2,3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := l1.String(); got != tt.want {
				t.Errorf("ListNode.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
