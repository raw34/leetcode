package _0034

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	binarySearch := func(nums []int, target int, lower bool) int {
		index := len(nums)
		left := 0
		right := len(nums) - 1
		for left <= right {
			mid := (right + left) / 2
			if nums[mid] > target || (nums[mid] == target && lower) {
				right = mid - 1
				index = mid
			} else {
				left = mid + 1
			}
		}

		return index
	}

	leftIdx := binarySearch(nums, target, true)
	rightIdx := binarySearch(nums, target, false) - 1
	if leftIdx <= rightIdx && rightIdx < len(nums) && nums[leftIdx] == target && nums[rightIdx] == target {
		res = []int{leftIdx, rightIdx}
	}

	return res
}
