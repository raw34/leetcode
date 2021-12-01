package _0033

func search(nums []int, target int) int {
	if nums == nil {
		return -1
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		val := nums[mid]

		if target == val {
			return mid
		}
		if nums[0] <= val {
			if nums[0] <= target && target < val {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[0] > target && target >= val {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
