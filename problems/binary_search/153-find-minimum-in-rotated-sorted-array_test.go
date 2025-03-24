package binary_search

import "testing"

func TestFindMin(t *testing.T) {
	numsList := [][]int{
		{3, 4, 5, 1, 2},
		{4, 5, 6, 7, 0, 1, 2},
		{11, 13, 15, 17},
	}
	for i := 0; i < len(numsList); i++ {
		t.Logf("input: %+v, output: %d", numsList[i], findMin(numsList[i]))
	}
}
