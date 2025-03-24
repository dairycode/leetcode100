package binary_search

import "testing"

func TestSearchRange(t *testing.T) {
	numsList := [][]int{
		{5, 7, 7, 8, 8, 10},
		{5, 7, 7, 8, 8, 10},
		{},
	}
	targetList := []int{8, 6, 0}
	for i := 0; i < len(numsList); i++ {
		nums, target := numsList[i], targetList[i]
		t.Logf("input: %+v %d, output: %+v", nums, target, searchRange(nums, target))
	}
}
