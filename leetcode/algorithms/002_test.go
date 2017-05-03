package algorithms

import (
	"testing"
	"fmt"
	"bytes"
	"strconv"
)

/**

https://leetcode.com/problems/add-two-numbers/#/description

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8

简单来说就是342+465=897
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	jinwei := false
	ret := &ListNode{0, nil}
	curRet := ret
	l1Point, l2Point := l1, l2
	for l1Point != nil && l2Point != nil {
		sum := l1Point.Val + l2Point.Val
		if jinwei {
			sum++
		}
		curRet.Next = &ListNode{sum % 10, nil}
		curRet = curRet.Next
		if sum >= 10 {
			jinwei = true
		} else {
			jinwei = false
		}
		l1Point = l1Point.Next
		l2Point = l2Point.Next
	}

	if l1Point != nil {
		curRet.Next = l1Point
		if jinwei {
			l1Point.Val++
			for l1Point.Next != nil && l1Point.Val >= 10 {
				l1Point.Val = 0
				l1Point = l1Point.Next
				l1Point.Val++
			}
			if l1Point.Val >= 10 {
				l1Point.Val = 0
				l1Point.Next = &ListNode{1, nil}
			}
		}
	} else if l2Point != nil {
		curRet.Next = l2Point
		if jinwei {
			l2Point.Val++
			for l2Point.Next != nil && l2Point.Val >= 10 {
				l2Point.Val = 0
				l2Point = l2Point.Next
				l2Point.Val++
			}
			if l2Point.Val >= 10 {
				l2Point.Val = 0
				l2Point.Next = &ListNode{1, nil}
			}
		}
	} else if jinwei {
		curRet.Next = &ListNode{1, nil}
	}

	return ret.Next
}

func TestAddTwoNumbers(t *testing.T) {
	testCases := [][][]int{
		{{2, 4, 3}, {5, 6, 4}, {7, 0, 8}},
		{{2, 4, 3, 1, 4}, {5, 6, 4}, {7, 0, 8, 1, 4}},
		{{2, 4, 3}, {5, 6, 4, 5, 9}, {7, 0, 8, 5, 9}},
		{{2, 4, 5}, {5, 6, 4}, {7, 0, 0, 1}},
		{{2, 4, 5, 9, 9}, {5, 6, 4}, {7, 0, 0, 0, 0, 1}},
		{
			{2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 9},
			{5, 6, 4, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 9, 9, 9, 9},
			{7, 0, 8, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 1, 4, 3, 9, 1},
		},
	}
	for _, testCase := range testCases {
		l1 := intAry2ListNode(testCase[0])
		l2 := intAry2ListNode(testCase[1])
		ret := addTwoNumbers(l1, l2)
		cur := ret
		for _, rint := range testCase[2] {
			if cur == nil || cur.Val != rint {
				t.Fatalf("结果不正确, 输入(%s) + (%s), 预期 %s, 实际 %s", formatListNode(l1), formatListNode(l2), formatIntAry(testCase[2]), formatListNode(ret))
			}
			cur = cur.Next
		}
		if cur != nil {
			t.Fatalf("结果长度不正确, 输入(%s) + (%s), 预期 %s, 实际 %s", formatListNode(l1), formatListNode(l2), formatIntAry(testCase[2]), formatListNode(ret))
		}
	}
}

func int2List(data int) *ListNode {
	ret := &ListNode{data % 10, nil}
	cur := ret
	data /= 10
	for data > 0 {
		cur.Next = &ListNode{data % 10, nil}
		cur = cur.Next
		data /= 10
	}
	return ret
}

func list2Int(node *ListNode) int {
	sum, bit := 0, 1
	for cur := node; cur != nil; cur = cur.Next {
		sum += cur.Val * bit
		bit *= 10
	}
	return sum
}

// 这个方法不适用于长数字
func _addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1int, l2int := list2Int(l1), list2Int(l2)
	sum := l1int + l2int
	return int2List(sum)
}

func formatIntAry(ary []int) string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	for idx, val := range ary {
		if idx > 0 {
			buffer.WriteString(fmt.Sprintf(", %d", val))
		} else {
			buffer.WriteString(strconv.Itoa(val))
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}

func formatListNode(node *ListNode) string {
	if node == nil {
		return ""
	}
	var buffer bytes.Buffer
	buffer.WriteString(strconv.Itoa(node.Val))
	for cur := node.Next; cur != nil; cur = cur.Next {
		buffer.WriteString(fmt.Sprintf(" -> %d", cur.Val))
	}
	return buffer.String()
}

func intAry2ListNode(ary []int) *ListNode {
	root := &ListNode{0, nil};
	cur := root
	for _, val := range ary {
		cur.Next = &ListNode{val, nil}
		cur = cur.Next
	}
	return root.Next
}
