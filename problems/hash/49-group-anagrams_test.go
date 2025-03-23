package hash

import (
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strsList := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
		{""},
		{"a"},
	}
	for _, strs := range strsList {
		ans := groupAnagrams(strs)
		t.Logf("input: %+v, output: %+v", strs, ans)
	}
}
