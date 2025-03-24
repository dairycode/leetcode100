package binary_search

import "testing"

func TestSearchInsert(t *testing.T) {
	numsList := [][]int{
		{1, 3, 5, 6},
		{1, 3, 5, 6},
		{1, 3, 5, 6},
	}
	targetList := []int{5, 2, 7}
	for i := 0; i < len(numsList); i++ {
		nums, target := numsList[i], targetList[i]
		t.Logf("input: %+v %d, output: %+v", nums, target, searchInsert(nums, target))
	}
}
