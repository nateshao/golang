package main

// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译。
// 本代码的正确性已通过力扣验证，如有疑问，可以对照 java 代码查看。

func lengthOfLongestSubstring(s string) int {
	window := make(map[rune]int)
	left, right := 0, 0
	// 记录结果
	res := 0
	for right < len(s) {
		c := rune(s[right])
		right++
		// 进行窗口内数据的一系列更新
		window[c] = window[c] + 1
		// 判断左侧窗口是否要收缩
		for window[c] > 1 {
			d := rune(s[left])
			left++
			// 进行窗口内数据的一系列更新
			window[d] = window[d] - 1
		}
		// 在这里更新答案
		if res < right-left {
			res = right - left
		}
	}
	return res
}

func main() {
	s := "abcabcbb"
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	slice = append(slice, 5)
	println(lengthOfLongestSubstring(s))
}

// 详细解析参见：
// https://labuladong.online/algo/essential-technique/sliding-window-framework/
