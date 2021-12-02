package _0154

func findMin(nums []int) int {
    length := len(nums)
    left := 0
    right := length - 1

    for left < right {
        mid := (left + right) / 2
        val := nums[mid]
        if val < nums[right] {
            right = mid
        } else if val > nums[right] {
            left = mid + 1
        } else {
            right--
        }
    }

    return nums[left]
}
