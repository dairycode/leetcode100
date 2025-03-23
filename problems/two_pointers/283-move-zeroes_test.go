package two_pointers

import "testing"

func TestMoveZeroes(t *testing.T) {
	numsList := [][]int{
		{0, 1, 0, 3, 12},
		{0},
	}
	for _, nums := range numsList {
		ans := append([]int{}, nums...)
		moveZeroes(ans)
		t.Logf("input: %+v, output: %+v", nums, ans)
	}
}
