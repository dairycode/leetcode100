package binary_search

import "testing"

func TestSearchMatrix(t *testing.T) {
	matrixList := [][][]int{
		{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
	}
	targetList := []int{3, 13}
	for i := 0; i < len(matrixList); i++ {
		matrix, target := matrixList[i], targetList[i]
		t.Logf("input: %+v %d, output: %t", matrix, target, searchMatrix(matrix, target))
	}
}
