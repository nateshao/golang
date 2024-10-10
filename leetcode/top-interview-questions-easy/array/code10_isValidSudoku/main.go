package main

import "fmt"

/*
*
有效的数独
请你判断一个9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字1-9在每一行只能出现一次。
数字1-9在每一列只能出现一次。
数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）

作者：LeetCode
链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2f9gg/
*/
func isValidSudoku(board [][]byte) bool {
	// 定义三个哈希表用于记录行、列、以及3x3宫格中的出现数字
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	// 初始化哈希表
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	// 遍历整个数独
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]

			// 忽略空格
			if num == '.' {
				continue
			}

			// 检查该数字是否已在当前行、列或3x3宫格中出现
			boxIndex := (i/3)*3 + j/3
			if rows[i][num] || cols[j][num] || boxes[boxIndex][num] {
				return false
			}

			// 将数字标记为已经出现
			rows[i][num] = true
			cols[j][num] = true
			boxes[boxIndex][num] = true
		}
	}
	return true
}
func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(board)) // 输出: true

}
