package sliding_window

// 找到字符串中所有字母异位词
// https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/
// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
// 示例：
// 输入: s = "cbaebabacd", p = "abc"
// 输出: [0,6]

func findAnagrams(s string, p string) []int {

	if len(s) < len(p) {
		return []int{}
	}

	ans := make([]int, 0)

	template := [26]int{}
	feature := [26]int{}

	for i := 0; i < len(p); i++ {
		template[p[i]-'a']++
		feature[s[i]-'a']++
	}
	if feature == template {
		ans = append(ans, 0)
	}

	left, right := 1, len(p)
	for right < len(s) {
		feature[s[left-1]-'a']--
		feature[s[right]-'a']++
		if feature == template {
			ans = append(ans, left)
		}
		left++
		right++
	}

	return ans
}
