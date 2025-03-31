package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "baissbaidubbbaidu"
	b := "baidu"
	c := "baissbb"

	// 将 a 中的 b 替换为空字符串
	result := strings.Replace(a, b, "", -1)

	// 输出结果
	fmt.Println("原始字符串 a:", a)
	fmt.Println("去掉 b 后的字符串:", result)
	fmt.Println("预期结果 c:", c)

	// 检查是否与 c 相等
	if result == c {
		fmt.Println("结果正确！")
	} else {
		fmt.Println("结果不正确！")
	}
}
