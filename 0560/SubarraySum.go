package _0560

func subarraySum(nums []int, k int) int {
    n := len(nums)
    preSum := make([]int, n+1)
    for i := 1; i <= n; i++ {
        preSum[i] = preSum[i-1] + nums[i-1]
    }

    res := 0
    for i := 1; i <= n; i++ {
        for j := 0; j < i; j++ {
            if preSum[i]-preSum[j] == k {
                res++
            }
        }
    }

    return res
}
