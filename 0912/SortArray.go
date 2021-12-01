package _912

import (
	"math"
)

/*
冒泡排序是一种简单的排序算法。
它重复地走访过要排序的数列，一次比较两个元素，如果它们的顺序错误就把它们交换过来。
走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。
这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。
*/

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 1; j < n-i; j++ {
			if arr[j] < arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}

	return arr
}

/*
选择排序(Selection-sort)是一种简单直观的排序算法。
它的工作原理：首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
以此类推，直到所有元素均排序完毕。
*/
func selectionSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		temp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp
	}

	return arr
}

/*
插入排序（Insertion-Sort）的算法描述是一种简单直观的排序算法。
它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
*/
func insertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		prevIndex := i - 1
		curr := arr[i]
		for prevIndex >= 0 && arr[prevIndex] > curr {
			arr[prevIndex+1] = arr[prevIndex]
			prevIndex--
		}
		arr[prevIndex+1] = curr
	}

	return arr
}

/*
1959年Shell发明，第一个突破O(n2)的排序算法，是简单插入排序的改进版。
它与插入排序的不同之处在于，它会优先比较距离较远的元素。希尔排序又叫缩小增量排序。
*/
func shellSort(arr []int) []int {
	return arr
}

/*
归并排序是建立在归并操作上的一种有效的排序算法。
该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。
将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。
若将两个有序表合并成一个有序表，称为2-路归并。
*/

func mergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}

	merge := func(left, right []int) []int {
		res := make([]int, 0)

		for len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				res = append(res, left[0])
				left = left[1:]
			} else {
				res = append(res, right[0])
				right = right[1:]
			}
		}

		for len(left) > 0 {
			res = append(res, left[0])
			left = left[1:]
		}

		for len(right) > 0 {
			res = append(res, right[0])
			right = right[1:]
		}

		return res
	}

	mid := int(math.Floor(float64(n / 2)))
	left := arr[0:mid]
	right := arr[mid:]

	return merge(mergeSort(left), mergeSort(right))
}

/*
快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。
*/
func quickSort(arr []int) []int {
	swap := func(arr []int, i, j int) []int {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp

		return arr
	}

	partition := func(arr []int, left, right int) int {
		pivot := left
		index := pivot + 1

		for i := index; i <= right; i++ {
			if arr[i] < arr[pivot] {
				swap(arr, i, index)
				index++
			}
		}
		swap(arr, pivot, index-1)

		return index - 1
	}

	var quick func(arr []int, left, right int) []int
	quick = func(arr []int, left, right int) []int {
		if left >= right {
			return arr
		}

		p := partition(arr, left, right)
		quick(arr, left, p-1)
		quick(arr, p+1, right)

		return arr
	}
	quick(arr, 0, len(arr)-1)

	return arr
}

/*
堆排序的思想就是先将待排序的序列建成大根堆，使得每个父节点的元素大于等于它的子节点。
此时整个序列最大值即为堆顶元素，我们将其与末尾元素交换，使末尾元素为最大值，然后再调整堆顶元素使得剩下的 n−1 个元素仍为大根堆，再重复执行以上操作我们即能得到一个有序的序列。
*/
func heapSort(arr []int) []int {
	length := len(arr)
	swap := func(arr []int, i, j int) []int {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp

		return arr
	}

	var heapUp func(arr []int, i int)
	heapUp = func(arr []int, i int) {
		left := 2*i + 1
		right := 2*i + 2
		max := i

		if left < length && arr[left] > arr[max] {
			max = left
		}

		if right < length && arr[right] > arr[max] {
			max = right
		}

		if max != i {
			swap(arr, i, max)
			heapUp(arr, max)
		}
	}

	buildMaxHeap := func(arr []int) {
		mid := int(math.Floor(float64(length / 2)))
		for i := mid; i >= 0; i-- {
			heapUp(arr, i)
		}
	}
	buildMaxHeap(arr)

	for i := len(arr) - 1; i >= 0; i-- {
		swap(arr, 0, i)
		length--
		heapUp(arr, 0)
	}

	return arr
}

/*
计数排序不是基于比较的排序算法，其核心在于将输入的数据值转化为键存储在额外开辟的数组空间中。
作为一种线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数。
*/
func countingSort(arr []int) []int {
	max := math.MinInt32
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	bucket := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]]++
	}

	sortedIndex := 0
	for i := 0; i < len(bucket); i++ {
		for bucket[i] > 0 {
			arr[sortedIndex] = i
			bucket[i]--
			sortedIndex++
		}
	}

	return arr
}

/*
桶排序是计数排序的升级版。
它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。
桶排序 (Bucket sort)的工作的原理：假设输入数据服从均匀分布，将数据分到有限数量的桶里，每个桶再分别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排）。
*/
func bucketSort(arr []int) []int {

	return arr
}

/*
基数排序是按照低位先排序，然后收集；再按照高位排序，然后再收集；依次类推，直到最高位。
有时候有些属性是有优先级顺序的，先按低优先级排序，再按高优先级排序。
最后的次序就是高优先级高的在前，高优先级相同的低优先级高的在前。
*/
func radixSort(arr []int) []int {

	return arr
}
