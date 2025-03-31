package main

import (
	"fmt"
)

func main() {
	a := "baissbaidubbbaidu"
	b := "baidu"
	//c := "baissbb"
	// 去掉 a 中的 b
	result := removeSubstring(a, b)
	fmt.Println(result)
}

// removeSubstring 去掉字符串 a 中的子字符串 b
func removeSubstring(a, b string) string {
	result := ""
	n := len(b)
	for i := 0; i < len(a); {
		// 检查是否匹配子字符串 b
		if i+n <= len(a) && a[i:i+n] == b {
			// 如果匹配，跳过子字符串 b
			i += n
		} else {
			// 如果不匹配，将当前字符加入结果
			result += string(a[i])
			i++
		}
	}
	return result
}
