package algorithms

import "testing"

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
	count := count1 + count2
	midIdx := count / 2
	start1, start2, end1, end2 := 0, 0, count1-1, count2-1
	mid := -1
	var midNum float64
	for start1 <= end1 {
		mid1 := (start1 + end1) / 2
		inAry2 := locaIntInAry(nums2, nums1[mid1], start2, end2)
		mid = mid1 + inAry2 + 1
		if mid == midIdx {
			midNum = float64(nums1[mid1])
			break
		} else if mid < midIdx {
			start1 = mid1 + 1
			start2 = inAry2
		} else {
			end1 = mid1 - 1
			end2 = inAry2
		}
	}
	if mid != midIdx {
		for start2 <= end2 {
			mid1 := (start2 + end2) / 2
			inAry2 := locaIntInAry(nums1, nums2[mid1], start1, end1)
			mid = mid1 + inAry2 + 1
			if mid == midIdx {
				midNum = float64(nums2[mid1])
				break
			} else if mid < midIdx {
				start1 = mid1 + 1
				start2 = inAry2
			} else {
				end1 = mid1 - 1
				end2 = inAry2
			}
		}
	}
	return midNum
}

/**
在start和end这2个下标之间，存在m个小于find的值，k个等于find的值，返回值n>=m+start且n<=m+k+start
 */
func locaIntInAry(nums []int, find, start, end int) int {
	for end-start > 1 {
		mid := (start + end) / 2
		if nums[mid] == find {
			return mid
		} else if nums[mid] > find {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if start == end {
		if nums[start] < find {
			return start + 1
		} else {
			return start
		}
	} else if nums[start] >= find {
		return start
	} else if nums[end] >= find {
		return end
	} else {
		return end + 1
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	type tmpS struct {
		ary  []int
		find int
	}
	list := []tmpS{
		{[]int{1, 3, 4, 8, 9}, 5},
		{[]int{1, 3, 4, 5, 5, 5, 5, 5, 5, 8, 9}, 5},
	}
	for _, tmp := range list {
		n := locaIntInAry(tmp.ary, tmp.find, 0, len(tmp.ary)-1)
		t.Logf("input : %s, find : %d, output : %d\n", formatIntAry(tmp.ary), tmp.find, n)
	}

	// return

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
