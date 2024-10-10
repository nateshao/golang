package main

// 旋转矩阵，将矩阵顺时针旋转90度
func rotate(matrix [][]int) {
	length := len(matrix)

	// 先上下交换
	for i := 0; i < length/2; i++ {
		matrix[i], matrix[length-i-1] = matrix[length-i-1], matrix[i]
	}

	// 再按对角线交换
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func main() {

}
