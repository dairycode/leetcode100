package hash

// 最长连续序列
// https://leetcode.cn/problems/longest-consecutive-sequence/description/
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 示例：
// 输入：nums = [100,4,200,1,3,2]
// 输出：4

func longestConsecutive(nums []int) int {

	m := make(map[int]struct{})

	for _, num := range nums {
		m[num] = struct{}{}
	}

	ans := 0

	for num := range m {
		if _, ok := m[num-1]; ok {
			continue
		}
		tmp := 1
		for {
			num++
			if _, ok := m[num]; !ok {
				break
			}
			tmp++
		}
		if tmp > ans {
			ans = tmp
		}
	}
	return ans
}
