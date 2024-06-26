package main

import (
	"fmt"
	"unicode/utf8"
)

//// map 错误示例
//func main() {
//	var m map[string]int
//	m["one"] = 1        // error: panic: assignment to entry in nil map
//	// m := make(map[string]int)// map 的正确声明，分配了实际的内存
//}

//// slice 正确示例
//func main() {
//	var s []int
//	s = append(s, 1)
//}

//// 数组使用值拷贝传参
//func main() {
//	x := [3]int{1, 2, 3}
//
//	func(arr [3]int) {
//		arr[0] = 7
//		fmt.Println(arr) // [7 2 3]
//	}(x)
//	fmt.Println(x) // [1 2 3]    // 并不是你以为的 [7 2 3]
//}

//// 传址会修改原数据
//func main() {
//	x := [3]int{1, 2, 3}
//
//	func(arr *[3]int) {
//		(*arr)[0] = 7
//		fmt.Println(arr) // &[7 2 3]
//	}(&x)
//	fmt.Println(x) // [7 2 3]
//}

//func main() {
//	x := []string{"a", "b", "c", "d", "e", "f", "g"}
//	for _, s := range x {
//		fmt.Println(s)
//	}
//}

// 使用各自独立的 6 个 slice 来创建 [2][3] 的动态多维数组
//func main() {
//	x := 2
//	y := 4
//
//	table := make([][]int, x)
//	for i := range table {
//		table[i] = make([]int, y)
//	}
//}

//func main() {
//	h, w := 2, 4
//	raw := make([]int, h*w)
//
//	for i := range raw {
//		raw[i] = i
//	}
//
//	// 初始化原始 slice
//	fmt.Println(raw, &raw[4]) // [0 1 2 3 4 5 6 7] 0xc420012120
//
//	table := make([][]int, h)
//	for i := range table {
//
//		// 等间距切割原始 slice，创建动态多维数组 table
//		// 0: raw[0*4: 0*4 + 4]
//		// 1: raw[1*4: 1*4 + 4]
//		table[i] = raw[i*w : i*w+w]
//	}
//
//	fmt.Println(table, &table[1][0]) // [[0 1 2 3] [4 5 6 7]] 0xc420012120
//}

func main() {
	char := "♥"
	fmt.Println(utf8.RuneCountInString(char)) // 1
}
