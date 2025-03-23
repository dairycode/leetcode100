package hash

import "testing"

func TestLongestConsecutive(t *testing.T) {
	numsList := [][]int{
		{100, 4, 200, 1, 3, 2},
		{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
		{1, 0, 1, 2},
	}
	for _, nums := range numsList {
		t.Logf("input: %+v, output: %+v", nums, longestConsecutive(nums))
	}
}
