package _0503

func nextGreaterElements(nums []int) []int {
    n := len(nums)
    nums = append(nums, nums...)
    stack := make([]int, 0)
    nextNums := map[int]int{}
    for i := 0; i < n*2; i++ {
        for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
            prev := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            nextNums[prev] = nums[i]
        }

        stack = append(stack, i)
        nextNums[i] = -1
    }

    res := make([]int, n)
    for i := 0; i < n; i++ {
        res[i] = nextNums[i]
    }

    return res
}
