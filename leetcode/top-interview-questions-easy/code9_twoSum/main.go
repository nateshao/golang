package main

import "fmt"

/**
public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            if (map.containsKey(target - nums[i])) {
                return new int[]{map.get(target - nums[i]), i};
            }
            map.put(nums[i], i);
        }
        return new int[]{0, 0};
    }
*/

func twoSum(nums []int, target int) []int {
	// 创建一个哈希表用于存储数值及其对应的索引
	hashMap := make(map[int]int)

	// 遍历数组
	for i, num := range nums {
		// 计算当前数 num 需要的目标补数
		complement := target - num

		// 如果补数存在于哈希表中，则找到了答案
		if index, found := hashMap[complement]; found {
			return []int{index, i}
		}

		// 否则，将当前数和它的索引加入哈希表
		hashMap[num] = i
	}

	// 如果没有找到，则返回空数组（题目保证了必定有解）
	return nil
}

func main() {
	// 示例 1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	result1 := twoSum(nums1, target1)
	fmt.Println(result1) // 输出: [0, 1]

	// 示例 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := twoSum(nums2, target2)
	fmt.Println(result2) // 输出: [1, 2]

	// 示例 3
	nums3 := []int{3, 3}
	target3 := 6
	result3 := twoSum(nums3, target3)
	fmt.Println(result3) // 输出: [0, 1]
}
