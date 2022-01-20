package _0303

type NumArray struct {
    preSum []int
}

func Constructor(nums []int) NumArray {
    n := len(nums)
    sum := make([]int, n+1)
    for i := 1; i <= n; i++ {
        sum[i] = sum[i-1] + nums[i-1]
    }

    return NumArray{preSum: sum}
}

func (this *NumArray) SumRange(left int, right int) int {
    return this.preSum[right+1] - this.preSum[left]
}
