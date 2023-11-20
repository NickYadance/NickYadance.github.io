package codetop

// TODO

import "testing"

func spiral(matrix [][]int, i, j int, path []int, record [][]bool) []int {
	path = append(path, matrix[i][j])
	record[i][j] = true
	if j+1 < len(matrix[0]) && !record[i][j+1] {
		return spiral(matrix, i, j+1, path, record)
	} else if i+1 < len(matrix) && !record[i+1][j] {
		return spiral(matrix, i+1, j, path, record)
	} else if j-1 >= 0 && !record[i][j-1] {
		return spiral(matrix, i, j-1, path, record)
	} else if i-1 >= 0 && !record[i-1][j] {
		return spiral(matrix, i-1, j, path, record)
	}
	return path
}

func spiralOrder(matrix [][]int) []int {
	record := make([][]bool, len(matrix))
	for i := range record {
		record[i] = make([]bool, len(matrix[0]))
	}
	return spiral(matrix, 0, 0, []int{}, record)
}

func Test54(t *testing.T) {
	spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}})
}
