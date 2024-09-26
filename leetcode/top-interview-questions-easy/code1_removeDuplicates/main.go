package main

import (
	"fmt"
	"testing"
)

/*
*
26. 删除有序数组中的重复项：https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：

更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k
*/
func removeDuplicates(nums []int) int {
	// 检查数组是否为空。如果数组长度为 0，返回 0，表示去重后没有元素。
	if len(nums) == 0 {
		return 0
	}

	// 初始化 left 指针为 0，表示当前不重复元素的最后位置。
	left := 0
	// 使用 right 指针从数组的第二个元素开始遍历（right = 1），直到数组末尾。
	for right := 1; right < len(nums); right++ {
		// 检查 left 指针指向的元素与 right 指针指向的元素是否不同。
		if nums[left] != nums[right] {
			// 如果不同，移动 left 指针到下一个位置。
			left++
			// 将 right 指向的唯一元素赋值给 nums[left]。
			nums[left] = nums[right]
		}
	}
	// 返回去重后数组的长度，left 是从 0 开始计数的，因此返回 left + 1。
	return left + 1
}

func TestRemoveDuplicates(t *testing.T) {
	// 示例 1
	nums1 := []int{1, 1, 2}
	expectedLength1 := 2
	expectedNums1 := []int{1, 2}

	length1 := removeDuplicates(nums1)

	if length1 != expectedLength1 {
		t.Errorf("expected length %d, got %d", expectedLength1, length1)
	}
	for i := 0; i < length1; i++ {
		if nums1[i] != expectedNums1[i] {
			t.Errorf("expected nums[%d] = %d, got %d", i, expectedNums1[i], nums1[i])
		}
	}

	// 示例 2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expectedLength2 := 5
	expectedNums2 := []int{0, 1, 2, 3, 4}

	length2 := removeDuplicates(nums2)

	if length2 != expectedLength2 {
		t.Errorf("expected length %d, got %d", expectedLength2, length2)
	}
	for i := 0; i < length2; i++ {
		if nums2[i] != expectedNums2[i] {
			t.Errorf("expected nums[%d] = %d, got %d", i, expectedNums2[i], nums2[i])
		}
	}
}

func main() {
	// 运行测试
	TestRemoveDuplicates(&testing.T{})
	fmt.Println("Tests completed.")
}
