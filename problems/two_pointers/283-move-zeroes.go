package two_pointers

// 移动零
// https://leetcode.cn/problems/move-zeroes/description/
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//
// 示例：
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]

func moveZeroes(nums []int) {

	left, right := 0, 0

	for right < len(nums) {

		if nums[right] == 0 {
			right++
			continue
		}

		nums[left] = nums[right]
		left++
		right++

	}

	for left < len(nums) {
		nums[left] = 0
		left++
	}

}
