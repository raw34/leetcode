package _0189

func rotate(nums []int, k int) {
    n := len(nums)
    temp := make([]int, n)
    for i := 0; i < n; i++ {
        j := (i + k) % n
        temp[j] = nums[i]
    }
    copy(nums, temp)
}
