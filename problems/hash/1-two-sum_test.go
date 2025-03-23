package hash

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	numsList := [][]int{
		{2, 7, 11, 15},
		{3, 2, 4},
		{3, 3},
	}
	targetList := []int{
		9,
		6,
		6,
	}
	for i := 0; i < len(numsList); i++ {
		nums, target := numsList[i], targetList[i]
		ans := twoSum(nums, target)
		t.Logf("input: %+v %d, output: %+v", nums, target, ans)
	}
}
