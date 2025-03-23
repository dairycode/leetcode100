package sliding_window

import "testing"

func TestFindAnagrams(t *testing.T) {
	sList := []string{
		"cbaebabacd",
		"abab",
	}
	pList := []string{
		"abc",
		"ab",
	}
	for i := 0; i < len(sList); i++ {
		s, p := sList[i], pList[i]
		t.Logf("input: %s %s, output: %+v", s, p, findAnagrams(s, p))
	}
}
