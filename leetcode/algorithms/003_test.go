package algorithms

import "testing"

/**

https://leetcode.com/problems/longest-substring-without-repeating-characters/#/description

Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */

func lengthOfLongestSubstring(s string) int {
	maxLen, curlen := 0, 0
	letterMap := make(map[int32]int)
	for idx, ch := range s {
		if val, ok := letterMap[ch]; ok {
			if curlen >= idx - val {
				if curlen > maxLen {
					maxLen = curlen
				}
				curlen = idx - val
			} else {
				curlen++
			}

		} else {
			curlen++
		}
		letterMap[ch] = idx
	}
	if curlen > maxLen {
		maxLen = curlen
	}
	return maxLen
}

func TestLengthOfLongestSubstring(t *testing.T) {
	type TestCases struct {
		input string
		len   int
	}
	testCases := []TestCases{
		{"", 0},
		{"abba", 2},
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, testCase := range testCases {
		if sublen := lengthOfLongestSubstring(testCase.input); sublen != testCase.len {
			t.Fatalf("输出结果不正确, 输入[%s], 预期输出[%d], 实际输出[%d]\n", testCase.input, testCase.len, sublen)
		}
	}
}
