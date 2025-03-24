package binary_search

// 搜索插入位置
// https://leetcode.cn/problems/search-insert-position/
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。
//
// 示例：
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2

func searchInsert(nums []int, target int) int {

	if len(nums) == 0 {
		return 0
	}

	ans := len(nums)

	left, right := 0, len(nums)-1

	for left <= right {

		mid := (left + right) / 2

		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return ans
}
