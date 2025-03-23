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
		t.Logf("input: %+v, output: %+v", strs, groupAnagrams(strs))
	}
}
