package main

import "fmt"

// 正确示例
func main() {
	x := []string{"a", "b", "c"}
	for _, v := range x { // 使用 _ 丢弃索引
		fmt.Println(v)
	}
}
