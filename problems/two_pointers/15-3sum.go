package two_pointers

import "sort"

// 三数之和
// https://leetcode.cn/problems/3sum/description/
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
//
// 示例：
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]

func threeSum(nums []int) [][]int {

	ans := make([][]int, 0)

	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}

		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left = left + 1; left < right; left++ {
					if nums[left] != nums[left-1] {
						break
					}
				}
				for right = right - 1; left < right; right-- {
					if nums[right] != nums[right+1] {
						break
					}
				}
			}
		}
	}

	return ans
}
