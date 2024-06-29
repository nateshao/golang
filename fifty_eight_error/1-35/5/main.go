package main

import "fmt"

// 正确示例
func main() {
	one := 0
	one, two := 1, 2    // two 是新变量，允许 one 的重复声明。比如 error 处理经常用同名变量 err
	one, two = two, one // 交换两个变量值的简写

	fmt.Println(one, two)
}
