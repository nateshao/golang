package main

/*
* 3. 无重复字符的最长子串 https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
 */
func lengthOfLongestSubstring(s string) int {
	windows := make(map[rune]int)
	left, right := 0, 0
	res := 0
	for right < len(s) {
		c := rune(s[right])
		right++
		windows[c]++
		for windows[c] > 1 {
			d := rune(s[left])
			left++
			windows[d]--
		}
		if res < right-left {
			res = right - left
		}
	}

	return res
}
