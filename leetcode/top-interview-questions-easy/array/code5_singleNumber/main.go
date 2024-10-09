package main

import "fmt"

/*
*1. 使用异或运算
利用异或运算的性质：相同的数异或结果为 0，任何数与 0 异或结果为该数本身。遍历整个数组，最终的结果就是只出现一次的数字。
*/
func singleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
	}
	return result
}

/*
*2. 使用哈希映射
使用哈希映射统计每个数字的出现次数，最后找到只出现一次的数字。虽然这种方法的时间复杂度仍为 O(n)，但空间复杂度为 O(n)。
*/
func singleNumber2(nums []int) int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	for k, v := range count {
		if v == 1 {
			return k
		}
	}
	return 0
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))       // 输出: 1
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2})) // 输出: 4
	fmt.Println(singleNumber([]int{1}))             // 输出: 1
}
