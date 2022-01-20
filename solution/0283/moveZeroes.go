package _0283

func moveZeroes(nums []int) {
    n := len(nums)
    for l := 0; l < n; l++ {
        left := nums[l]
        if left == 0 {
            for r := l + 1; r < n; r++ {
                right := nums[r]
                if right != 0 {
                    nums[l], nums[r] = nums[r], nums[l]
                    break
                }
            }
        }
    }
}

func moveZeroes2(nums []int) {
    n := len(nums)
    for left, right := 0, 0; right < n; right++ {
        if nums[right] != 0 {
            nums[left], nums[right] = nums[right], nums[left]
            left++
        }
    }
}
