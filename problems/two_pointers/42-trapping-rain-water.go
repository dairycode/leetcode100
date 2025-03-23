package two_pointers

// 接雨水
// https://leetcode.cn/problems/trapping-rain-water/description/
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
// 示例：
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6

func trap(height []int) int {

	if len(height) <= 2 {
		return 0
	}

	ans := 0
	leftMax, rightMax := height[0], height[len(height)-1]
	left, right := 1, len(height)-2

	for left <= right {
		if leftMax < rightMax {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
	}
	return ans
}
