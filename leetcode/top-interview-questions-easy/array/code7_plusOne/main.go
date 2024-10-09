package main

import "fmt"

func plusOne(digits []int) []int {
	// 从数组末尾开始处理加1操作
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++ // 当前位加1
		if digits[i] < 10 {
			return digits // 如果没有进位，直接返回结果
		}
		digits[i] = 0 // 如果进位，需要将当前位设为0
	}
	return append([]int{1}, digits...)
}

func main() {
	// 测试用例
	fmt.Println(plusOne([]int{1, 2, 3}))    // 输出: [1, 2, 4]
	fmt.Println(plusOne([]int{4, 3, 2, 1})) // 输出: [4, 3, 2, 2]
	fmt.Println(plusOne([]int{9, 9, 9}))    // 输出: [1, 0, 0, 0]
	fmt.Println(plusOne([]int{0}))          // 输出: [1]
}
