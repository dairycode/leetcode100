package two_pointers

import "testing"

func TestThreeSum(t *testing.T) {
	numsList := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{0, 1, 1},
		{0, 0, 0},
	}
	for _, nums := range numsList {
		t.Logf("input: %+v, output: %+v", nums, threeSum(nums))
	}
}
