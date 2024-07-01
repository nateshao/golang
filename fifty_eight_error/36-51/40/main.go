package main

import (
	"fmt"
)

// 正确的 recover 调用示例
func main() {
	defer func() {
		fmt.Println("recovered: ", recover())
	}()
	panic("not good")
}
