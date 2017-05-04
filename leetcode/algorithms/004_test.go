package algorithms

import (
	"testing"
)

/**
https://leetcode.com/problems/median-of-two-sorted-arrays/#/description

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5

 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	count1, count2 := len(nums1), len(nums2)
	if count1 > count2 {
		nums1, nums2, count1, count2 = nums2, nums1, count2, count1
	}
	if count2 == 0 {
		return 0
	}
	imin, imax, halfLen := 0, count1, (count1+count2+1)/2

	for imin <= imax {
		i := (imin + imax) / 2
		j := halfLen - i
		if i < count1 && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else {
			var maxLeft int
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else if nums1[i-1] > nums2[j-1] {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = nums2[j-1]
			}
			if (count1+count2)%2 == 1 {
				return float64(maxLeft)
			}
			var minRight int
			if i == count1 {
				minRight = nums2[j]
			} else if j == count2 {
				minRight = nums1[i]
			} else if nums1[i] < nums2[j] {
				minRight = nums1[i]
			} else {
				minRight = nums2[j]
			}
			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0
}

func TestFindMedianSortedArrays(t *testing.T) {
	type TestCases struct {
		nums1  []int
		nums2  []int
		median float64
	}
	testCases := []TestCases{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}
	for _, tc := range testCases {
		if median := findMedianSortedArrays(tc.nums1, tc.nums2); median != tc.median {
			t.Fatalf("结果不正确, 输入%s, %s, 预期%.1f, 输出%.1f", formatIntAry(tc.nums1), formatIntAry(tc.nums2), tc.median, median)
		}
	}
}
