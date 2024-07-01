package main

import "fmt"

// 如果要修改原有元素的值，应该使用索引直接访问：
//func main() {
//	data := []int{1, 2, 3}
//	for i, v := range data {
//		data[i] = v * 10
//	}
//	fmt.Println("data: ", data) // data:  [10 20 30]
//}

func main() {
	data := []*struct{ num int }{{1}, {2}, {3}}
	for _, v := range data {
		v.num *= 10 // 直接使用指针更新
	}
	fmt.Println(data[0], data[1], data[2]) // &{10} &{20} &{30}
}
