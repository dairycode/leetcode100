package binary_search

import "testing"

func TestSearch(t *testing.T) {
	numsList := [][]int{
		{4, 5, 6, 7, 0, 1, 2},
		{4, 5, 6, 7, 0, 1, 2},
		{1},
	}
	targetList := []int{0, 3, 0}
	for i := 0; i < len(numsList); i++ {
		nums, target := numsList[i], targetList[i]
		t.Logf("input: %+v %d, output: %d", nums, target, search(nums, target))
	}
}
