package algorithms

import (
	"testing"
)

/**
https://leetcode.com/problems/longest-palindromic-substring/#/description

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example:

Input: "babad"

Output: "bab"

Note: "aba" is also a valid answer.
Example:

Input: "cbbd"

Output: "bb"

一个长度不超过1000的字符串，找出其中最长的回文子串

 */

// 7.25%，太失败了
func _longestPalindrome(s string) string {
	strlen := len(s)
	if strlen <= 1 {
		return s
	} else if strlen == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	}
	list1 := make(map[int]bool)
	list2 := make(map[int]bool)
	for idx := 0; idx < strlen; idx++ {
		list1[idx] = true
		if idx+1 < strlen && s[idx] == s[idx+1] {
			list2[idx] = true
		}
	}
	maxLen := 1
	for len(list1) > 0 {
		list3 := make(map[int]bool)
		for idx := range list1 {
			if idx > 0 && idx+maxLen < strlen && s[idx-1] == s[idx+maxLen] {
				list3[idx-1] = true
			}
		}
		if len(list3) == 0 {
			if len(list2) == 0 {
				break
			} else {
				list1 = list2
				list2 = list3
				maxLen++
			}
		} else if len(list2) == 0 {
			list1 = list3
			maxLen += 2
		} else {
			list1 = list2
			list2 = list3
			maxLen++
		}
	}
	for idx := range list1 {
		jdx := idx + maxLen
		return s[idx:jdx]
	}
	return ""
}

// 最优解
func longestPalindrome(s string) string {
	answer := 0
	longestLen := 1
	for index := 0; index < len(s); {
		if len(s)-index < longestLen/2 {
			break
		}
		frontIndex, endIndex := index, index
		for endIndex < len(s)-1 && s[endIndex] == s[endIndex+1] {
			endIndex++
		}
		index = endIndex + 1
		for frontIndex > 0 && endIndex < len(s)-1 && s[frontIndex-1] == s[endIndex+1] {
			frontIndex--
			endIndex++
		}
		if endIndex-frontIndex+1 > longestLen {
			longestLen = endIndex - frontIndex + 1
			answer = frontIndex
		}
	}
	return s[answer: answer+longestLen]
}

// Manacher算法 http://blog.csdn.net/yzl_rex/article/details/7908259 88.42%
func longestPalindromeManacher(s string) string {
	strlen := len(s)
	if strlen <= 1 {
		return s
	}
	maxLen := 1
	left := 0
	manacherLen := 2*strlen + 1
	maxId, maxMid := 0, 0
	aryP := make([]int, manacherLen, manacherLen)
	for idx := 0; idx + maxLen < manacherLen; idx++ {
		if maxId > idx {
			if aryP[2*maxMid-idx] < maxId-idx {
				aryP[idx] = aryP[2*maxMid-idx]
			} else {
				aryP[idx] = maxId - idx
			}
		} else {
			aryP[idx] = 1
		}
		for idx-aryP[idx] >= 0 && idx+aryP[idx] < manacherLen && compStringManacher(s, idx+aryP[idx], idx-aryP[idx]) {
			aryP[idx]++
		}
		if aryP[idx] + idx > maxId {
			maxId = aryP[idx] + idx
			maxMid = idx
		}
		if maxLen < aryP[idx]-1 {
			maxLen = aryP[idx] - 1
			left = (idx - aryP[idx] + 1) / 2
		}
	}
	return s[left:left+maxLen]
}

func compStringManacher(s string, idx1, idx2 int) bool {
	mod1, mod2 := idx1%2, idx2%2
	if mod1 != mod2 {
		return false
	} else if mod1 == 0 {
		return true
	} else {
		return s[(idx1-1)/2] == s[(idx2-1)/2]
	}
}

// 14.49% 还是很差。。。而且不加第一个判断特殊处理还会超时！
func __longestPalindrome(s string) string {
	if isPalindromic(s) {
		return s
	}
	strlen := len(s)
	if strlen <= 1 {
		return s
	}
	letters := make(map[int32][]int)
	maxLen := 1
	maxLeft := 0
	subMap := make(map[int]map[int]bool, strlen)
	subMap[1] = make(map[int]bool, strlen)
	for idx, ch := range s {
		subMap[1][idx] = true
		if _, ok := letters[ch]; ok {
			for _, left := range letters[ch] {
				subLen := idx - left - 1
				if _, hasSub := subMap[subLen]; hasSub || subLen == 0 {
					if _, hasSub = subMap[subLen][left+1]; hasSub || subLen == 0 {
						if _, ok = subMap[subLen+2]; !ok {
							subMap[subLen+2] = make(map[int]bool, strlen)
							if subLen+2 > maxLen {
								maxLeft = left
								maxLen = subLen + 2
							}
						}
						subMap[subLen+2][left] = true
					}
				}
			}
		} else {
			letters[ch] = make([]int, 0, strlen);
		}
		letters[ch] = append(letters[ch], idx)
	}
	jdx := maxLen + maxLeft
	return s[maxLeft:jdx]
}

func isPalindromic(str string) bool {
	strlen := len(str)
	if strlen == 0 {
		return true
	}
	mid := strlen / 2
	for idx := 0; idx <= mid; idx++ {
		jdx := strlen - 1 - idx
		if str[idx] != str[jdx] {
			return false
		}
	}
	return true
}

func TestLongestPalindrome(t *testing.T) {
	type TestCases struct {
		input string
		len   int
	}
	testCases := []TestCases{
		{"", 0},
		{"cbbd", 2},
		{"abba", 4},
		{"babad", 3},
		{"zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz", 1000},
	}

	for _, testCase := range testCases {
		substr := longestPalindromeManacher(testCase.input)
		sublen := len(substr)
		if sublen != testCase.len {
			t.Fatalf("输出结果长度不正确, 输入[%s], 预期输出[%d], 实际输出[%d]\n", testCase.input, testCase.len, sublen)
		} else if !isPalindromic(substr) {
			t.Fatalf("输出结果不是正确的回文, 输入[%s], 输出[%s]\n", testCase.input, substr)
		}
	}
}
