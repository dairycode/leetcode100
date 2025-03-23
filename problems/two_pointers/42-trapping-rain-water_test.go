package two_pointers

import "testing"

func TestTrap(t *testing.T) {
	heightList := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{4, 2, 0, 3, 2, 5},
	}
	for _, height := range heightList {
		t.Logf("input: %+v, output: %d", height, trap(height))
	}
}
