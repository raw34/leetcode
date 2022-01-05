package _0496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    m := len(nums2)
    n := len(nums1)
    // 利用单调栈计算nums2中每个数字的下一个最大数字
    nums := map[int]int{}
    stack := make([]int, 0)
    for i := 0; i < m; i++ {
        curr := nums2[i]
        // 当栈不为空，且当前数字大于栈顶数字时，当前数字为栈顶数字的下一个最大数字
        for len(stack) > 0 && stack[len(stack)-1] < curr {
            prev := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            nums[prev] = curr
        }
        // 否则入栈当前数字
        stack = append(stack, curr)
        // 给每个入栈的数字赋默认值
        nums[curr] = -1
    }

    // 遍历nums1，从计算好的最大数字数组nums2中获取结果
    res := make([]int, n)
    for j := 0; j < n; j++ {
        res[j] = nums[nums1[j]]
    }

    return res
}
