package main

/*
*
把非0的往前挪
*/
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	// 将所有非零元素按顺序放置在数组前面
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	// 将剩余的位置填充为0
	for index < len(nums) {
		nums[index] = 0
		index++
	}
}

/*
*
4. 单次遍历进行交换
该方法类似双指针方法，使用一个变量来记录当前最左侧的 0 元素的位置，然后遍历数组时不断与非零元素交换位置。这样只需要一次遍历。
*/
func moveZeroesSwap(nums []int) {
	zeroIndex := 0 // 用于跟踪0的插入位置
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// 交换当前元素和 zeroIndex 位置的元素
			nums[zeroIndex], nums[i] = nums[i], nums[zeroIndex]
			zeroIndex++
		}
	}
}

func main() {

}
