package two_pointers

import "leetcode100/util"

// 盛水最多的容器
// https://leetcode.cn/problems/container-with-most-water/description/
// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 说明：你不能倾斜容器。
//
// 示例：
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49

func maxArea(height []int) int {

	ans := 0
	left, right := 0, len(height)-1

	for left < right {

		h := util.Min(height[left], height[right])
		area := h * (right - left)
		if area > ans {
			ans = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}

	return ans
}
