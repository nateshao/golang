package main

import "fmt"

/*
*
算法思路：
遍历整个股票交易日价格列表 price，并执行贪心策略：所有上涨交易日都买卖（赚到所有利润），所有下降交易日都不买卖（永不亏钱）。

1.设 tmp 为第 i-1 日买入与第 i 日卖出赚取的利润，即 tmp = prices[i] - prices[i - 1] ；
2.当该天利润为正 tmp > 0，则将利润加入总利润 profit；当利润为 0 或为负，则直接跳过；
3.遍历完成后，返回总利润 profit。

作者：Krahets
链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/solutions/12625/best-time-to-buy-and-sell-stock-ii-zhuan-hua-fa-ji/

复杂度分析：
时间复杂度 O(N) ： 只需遍历一次 price 。
空间复杂度 O(1) ： 变量使用常数额外空间。
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profit := 0
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			profit += tmp
		}
	}
	return profit
}

func main() {
	nums1 := []int{7, 1, 5, 3, 6, 4}
	nums2 := []int{1, 2, 3, 4, 5}
	nums3 := []int{7, 6, 4, 3, 1}
	profit1 := maxProfit(nums1)
	profit2 := maxProfit(nums2)
	profit3 := maxProfit(nums3)

	fmt.Printf("result1: %d, result2:  %d, result2:  %d\n", profit1, profit2, profit3)
}
