package algorithms

import "testing"

/**
	https://leetcode.com/problems/two-sum/#/description

	Given an array of integers, return indices of the two numbers such that they add up to a specific target.

	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	给一组整数和一个目标整数，找出2个数字相加等于目标数字，返回找出的2个数字下标

	数组中可能出现相同数字
 */

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	var (
		key   bool
		value int
	)

	ret := []int{0, 0}
	for i, v := range nums {
		if value, key = numMap[v]; key {
			ret[0] = i
			ret[1] = value
			return ret
		} else {
			numMap[target-v] = i
		}
	}
	return ret
}

func TestTwoSum(t *testing.T) {
	type TestCaseTwoSum struct {
		nums   []int
		target int
	}
	cases := []TestCaseTwoSum{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{2, 5, 14}, 7},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 2, 3}, 6},
	}
	for _, tc := range cases {
		index := twoSum(tc.nums, tc.target)
		if len(index) != 2 {
			t.Error("返回参数个数不正确, 输入", tc, ", 输出", index, "\n")
		}
		if index[0] == index[1] {
			t.Error("一个元素用了2次, 输入", tc, ", 输出", index, "\n")
		}
		numCount := len(tc.nums)
		for i, idx := range index {
			if idx < 0 || idx >= numCount {
				t.Error("第[", i, "]个输出不合法, 输入", tc, ", 输出", index, "\n")
			}
		}
		if tc.nums[index[0]] + tc.nums[index[1]] != tc.target {
			t.Error("相加的值不正确, 输入", tc, ", 输出", index, "\n")
		}
	}
}
