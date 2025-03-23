package sliding_window

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	sList := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
	}
	for _, s := range sList {
		t.Logf("input: %s, output:%d", s, lengthOfLongestSubstring(s))
	}
}
