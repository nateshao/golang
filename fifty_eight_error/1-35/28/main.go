package main

import (
	"fmt"
)

// 正确示例
func main() {
	var d uint8 = 2
	fmt.Printf("%08b\n", d)  // 00000010
	fmt.Printf("%08b\n", ^d) // 11111101
}
