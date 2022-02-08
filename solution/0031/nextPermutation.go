package _0031

/*
下一个排列总是比当前排列要大，除非该排列已经是最大的排列。我们希望找到一种方法，能够找到一个大于当前序列的新序列，且变大的幅度尽可能小。具体地：
1、我们需要将一个左边的「较小数」与一个右边的「较大数」交换，以能够让当前排列变大，从而得到下一个排列。
2、同时我们要让这个「较小数」尽量靠右，而「较大数」尽可能小。当交换完成后，「较大数」右边的数需要按照升序重新排列。这样可以在保证新排列大于原来排列的情况下，使变大的幅度尽可能小。
*/
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
