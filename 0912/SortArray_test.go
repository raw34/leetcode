package _912

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortArray(t *testing.T) {
	expend := []int{1, 2, 3, 5}

	arr := []int{5, 3, 1, 2}
	assert.Equal(t, expend, bubbleSort(arr))
	arr = []int{5, 3, 1, 2}
	assert.Equal(t, expend, selectionSort(arr))
	arr = []int{5, 3, 1, 2}
	assert.Equal(t, expend, insertionSort(arr))
	arr = []int{5, 3, 1, 2}
	assert.Equal(t, expend, mergeSort(arr))
	arr = []int{5, 3, 1, 2}
	assert.Equal(t, expend, quickSort(arr))
	arr = []int{5, 3, 1, 2}
	assert.Equal(t, expend, heapSort(arr))
}
