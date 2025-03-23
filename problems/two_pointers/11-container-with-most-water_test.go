package two_pointers

import "testing"

func TestMaxArea(t *testing.T) {
	heightList := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
	}
	for _, height := range heightList {
		t.Logf("input: %+v, output: %d", height, maxArea(height))
	}
}
