package binary_search

// 搜索二维矩阵
// https://leetcode.cn/problems/search-a-2d-matrix/description/
// 给你一个满足下述两条属性的 m x n 整数矩阵：
// 每行中的整数从左到右按非严格递增顺序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
//
// 示例：
// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// 输出：true

func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row := len(matrix) - 1

	left, right := 0, len(matrix)-1

	for left <= right {

		mid := (left + right) / 2
		if target == matrix[mid][0] {
			return true
		} else if target < matrix[mid][0] {
			right = mid - 1
		} else {
			row = mid
			left = mid + 1
		}

	}

	left, right = 0, len(matrix[row])-1

	for left <= right {

		mid := (left + right) / 2
		if target == matrix[row][mid] {
			return true
		} else if target < matrix[row][mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return false
}
