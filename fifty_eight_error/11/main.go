package main

import "fmt"

// 正确示例
func main() {
	var s string // 字符串类型的零值是空串 ""
	if s == "" {
		s = "default"
	}
	fmt.Println(s)
}
