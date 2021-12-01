package _912

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortArray(t *testing.T) {
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	arr := []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expect, bubbleSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expect, selectionSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expect, insertionSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expect, shellSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expect, mergeSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expect, quickSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expect, heapSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expect, countingSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expect, bucketSort(arr))
	arr = []int{5, 3, 1, 2, 6, 7, 10, 8, 4, 9}
	assert.Equal(t, expect, radixSort(arr))
}
