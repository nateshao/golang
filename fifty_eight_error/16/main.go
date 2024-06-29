package main

import "fmt"

// 修改示例
func main() {
	x := "text"
	xBytes := []byte(x)
	xBytes[0] = 'T' // 注意此时的 T 是 rune 类型
	x = string(xBytes)
	fmt.Println(x) // Text
}
