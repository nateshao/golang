# 【leetcode初级算法】122. 买卖股票的最佳时机 II
## 题目描述

给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

返回 你能获得的 最大 利润 。

示例 1：

输入：prices = [7,1,5,3,6,4]
输出：7
解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3。
最大总利润为 4 + 3 = 7 。
示例 2：

输入：prices = [1,2,3,4,5]
输出：4
解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
最大总利润为 4 。
示例 3：

输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0。

## 解题思路

### 具体思路
该函数的目的是通过**贪心算法**，在股票价格的上涨交易日进行买入和卖出，以获取最大利润。具体步骤如下：

#### 1. **遍历价格列表**
遍历整个股票交易日价格列表 `prices`，并执行贪心策略：所有上涨交易日都进行买卖（赚到所有利润），所有下降交易日都不买卖（避免亏损）。

#### 2. **利润计算**
- 设 `tmp` 为第 `i-1` 日买入与第 `i` 日卖出赚取的利润，即 `tmp = prices[i] - prices[i - 1]`。
- 当该天利润为正 `tmp > 0`，则将利润加入总利润 `profit`；如果利润为 0 或为负，则跳过。

#### 3. **返回结果**
- 遍历完成后，返回总利润 `profit`。

### 示例
假设输入价格数组为 `[7, 1, 5, 3, 6, 4]`：
- 第一天以 `1` 价格买入，第二天以 `5` 价格卖出，赚取 `4` 的利润。
- 第三天以 `3` 价格买入，第四天以 `6` 价格卖出，赚取 `3` 的利润。
- 返回的总利润为 `7`。

### 总结
该算法通过有效地使用一次遍历，确保了在时间复杂度 O(N) 内完成最大利润的计算，且空间复杂度为 O(1)，仅使用了常数的额外空间。

## Golang 实现

```go
package main

import "fmt"

// maxProfit 计算股票的最大利润
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

	fmt.Printf("result1: %d, result2: %d, result3: %d\n", profit1, profit2, profit3)
}
```

**复杂度分析**

- **时间复杂度**: O(N) —— 只需遍历一次 `prices` 数组。
- **空间复杂度**: O(1) —— 变量使用常数额外空间。

![](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/wximage-20240928165106435.png)

## Java 实现

```java
public class Solution {
    public int maxProfit(int[] prices) {
        // 检查数组是否为空。如果数组长度为 0，返回 0。
        if (prices.length == 0) {
            return 0;
        }

        int profit = 0;
        // 遍历价格数组，从第二个元素开始。
        for (int i = 1; i < prices.length; i++) {
            int tmp = prices[i] - prices[i - 1];
            // 如果利润为正，则加入总利润。
            if (tmp > 0) {
                profit += tmp;
            }
        }
        return profit;
    }
    
    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] prices1 = {7, 1, 5, 3, 6, 4};
        int[] prices2 = {1, 2, 3, 4, 5};
        int[] prices3 = {7, 6, 4, 3, 1};
        System.out.println("Profit 1: " + solution.maxProfit(prices1)); // 7
        System.out.println("Profit 2: " + solution.maxProfit(prices2)); // 4
        System.out.println("Profit 3: " + solution.maxProfit(prices3)); // 0
    }
}
```

## PHP 实现

```php
function maxProfit($prices) {
    // 检查数组是否为空。如果数组长度为 0，返回 0。
    if (count($prices) == 0) {
        return 0;
    }

    $profit = 0;
    // 遍历价格数组，从第二个元素开始。
    for ($i = 1; $i < count($prices); $i++) {
        $tmp = $prices[$i] - $prices[$i - 1];
        // 如果利润为正，则加入总利润。
        if ($tmp > 0) {
            $profit += $tmp;
        }
    }
    return $profit;
}

// 测试
$prices1 = [7, 1, 5, 3, 6, 4];
$prices2 = [1, 2, 3, 4, 5];
$prices3 = [7, 6, 4, 3, 1];
echo "Profit 1: " . maxProfit($prices1) . "\n"; // 7
echo "Profit 2: " . maxProfit($prices2) . "\n"; // 4
echo "Profit 3: " . maxProfit($prices3) . "\n"; // 0
```

## C++ 实现

```cpp
#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    int maxProfit(vector<int>& prices) {
        // 检查数组是否为空。如果数组长度为 0，返回 0。
        if (prices.empty()) {
            return 0;
        }

        int profit = 0;
        // 遍历价格数组，从第二个元素开始。
        for (int i = 1; i < prices.size(); i++) {
            int tmp = prices[i] - prices[i - 1];
            // 如果利润为正，则加入总利润。
            if (tmp > 0) {
                profit += tmp;
            }
        }
        return profit;
    }
};

int main() {
    Solution solution;
    vector<int> prices1 = {7, 1, 5, 3, 6, 4};
    vector<int> prices2 = {1, 2, 3, 4, 5};
    vector<int> prices3 = {7, 6, 4, 3, 1};
    cout << "Profit 1: " << solution.maxProfit(prices1) << endl; // 7
    cout << "Profit 2: " << solution.maxProfit(prices2) << endl; // 4
    cout << "Profit 3: " << solution.maxProfit(prices3) << endl; // 0
    return 0;
}
```