package _912

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortArray(t *testing.T) {
	expend := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	arr := []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expend, bubbleSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expend, selectionSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expend, insertionSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expend, mergeSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expend, quickSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expend, heapSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expend, countingSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expend, bucketSort(arr))
}
