package _912

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortArray(t *testing.T) {
	arr := []int{5, 3, 1, 2}
	expend := []int{1, 2, 3, 5}

	assert.Equal(t, expend, bubbleSort(arr))
	assert.Equal(t, expend, selectionSort(arr))
	assert.Equal(t, expend, insertionSort(arr))
	assert.Equal(t, expend, mergeSort(arr))
	assert.Equal(t, expend, quickSort(arr))
}
