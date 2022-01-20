package _0035

func searchInsert(nums []int, target int) int {
	res := len(nums)
	left := 0
	right := res - 1

	for left <= right {
		mid := (right-left)/2 + left
		val := nums[mid]
		if val >= target {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return res
}
