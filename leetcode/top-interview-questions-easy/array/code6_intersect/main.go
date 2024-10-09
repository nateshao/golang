package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2y0c2/
给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。

示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]
*/

/*
*解法一：排序 + 双指针
 */
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	result := []int{}
	i, j := 0, 0

	// 双指针遍历两个数组
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			result = append(result, nums1[i])
			i++
			j++
		}
	}

	return result
}

/*
*
方法 1：使用哈希映射
通过哈希映射记录每个元素的出现次数，然后遍历第二个数组，找到交集。
*/
func intersect2(nums1 []int, nums2 []int) []int {
	countMap := make(map[int]int)
	result := []int{}

	// 记录 nums1 中每个元素的出现次数
	for _, num := range nums1 {
		countMap[num]++
	}

	// 遍历 nums2，找到交集并更新计数
	for _, num := range nums2 {
		if count, ok := countMap[num]; ok && count > 0 {
			result = append(result, num)
			countMap[num]-- // 减少出现次数
		}
	}
	return result
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))       // 输出: [2, 2]
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})) // 输出: [4, 9]
}
