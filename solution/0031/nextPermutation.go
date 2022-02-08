package _0031

func nextPermutation(nums []int) {
    reverse := func(nums []int, i int, j int) {
        for i < j {
            nums[i], nums[j] = nums[j], nums[i]
            i++
            j--
        }
    }
    // 找到第一个逆序对
    i := len(nums) - 2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }
    // 如果没有逆序对，则按照降序排列
    if i < 0 {
        reverse(nums, 0, len(nums)-1)
        return
    }
    // 找到第一个比i大的数
    j := len(nums) - 1
    for j >= 0 && nums[j] <= nums[i] {
        j--
    }
    // 交换i和j
    nums[i], nums[j] = nums[j], nums[i]
    // 反转i+1到末尾
    reverse(nums, i+1, len(nums)-1)
}
