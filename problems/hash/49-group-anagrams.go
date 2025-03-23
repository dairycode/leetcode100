package hash

// 字母异位词分组
// https://leetcode.cn/problems/group-anagrams/description/
// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
//
// 示例：
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

func groupAnagrams(strs []string) [][]string {

	m := make(map[[26]int][]string)

	for _, str := range strs {

		key := [26]int{}
		for i := 0; i < len(str); i++ {
			key[str[i]-'a']++
		}
		m[key] = append(m[key], str)

	}

	ans := make([][]string, 0, len(m))

	for _, v := range m {
		ans = append(ans, v)
	}

	return ans
}
