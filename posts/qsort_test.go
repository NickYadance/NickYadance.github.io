package posts

import (
	"gotest.tools/assert"
	"math/rand"
	"testing"
)

func swap(A []int, i, j int) {
	tmp := A[i]
	A[i] = A[j]
	A[j] = tmp
}

func random(l, r int) int {
	return l + rand.Intn(r-l)
}

func Partition(A []int, l, r int) (pivot int) {
	i := l
	j := l
	x := A[r-1]
	for j < r {
		if A[j] <= x {
			swap(A, i, j)
			i = i + 1
		}
		j = j + 1
	}
	return i - 1
}

func QSort(A []int, l, r int) {
	if l < r {
		p := random(l, r)
		swap(A, p, r-1)
		pivot := Partition(A, l, r)
		QSort(A, l, pivot)
		QSort(A, pivot+1, r)
	}
}

func TestQSort(t *testing.T) {
	A := []int{2, 1, 9, 10, 0, 0}
	QSort(A, 0, len(A))
	assert.DeepEqual(t, A, []int{0, 0, 1, 2, 9, 10})
}
