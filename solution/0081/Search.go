package _0081

func search(nums []int, target int) bool {
	if nums == nil {
		return false
	}

	length := len(nums)
	left := 0
	right := length - 1
	for left <= right {
		mid := (left + right) / 2
		val := nums[mid]
		if target == val {
			return true
		}
		if nums[left] == val && nums[right] == val {
			left++
			right--
		} else if nums[left] <= val {
			if nums[left] <= target && target < val {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[length-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}
