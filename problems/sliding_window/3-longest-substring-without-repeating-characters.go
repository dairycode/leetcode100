package sliding_window

// 无重复字符的最长字串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
//
// 示例：
// 输入: s = "abcabcbb"
// 输出: 3

func lengthOfLongestSubstring(s string) int {

	ans := 0

	m := make(map[byte]struct{})
	left, right := 0, 0

	for right < len(s) {

		if _, ok := m[s[right]]; !ok {
			m[s[right]] = struct{}{}
			right++
			if right-left > ans {
				ans = right - left
			}
			continue
		}

		for {
			delete(m, s[left])
			left++
			if _, ok := m[s[right]]; !ok {
				break
			}
		}
	}
	return ans
}
