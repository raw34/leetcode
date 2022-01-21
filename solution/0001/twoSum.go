package _0001

func twoSum(nums []int, target int) []int {
    n := len(nums)
    hash := map[int]int{}
    for i := 0; i < n; i++ {
        hash[nums[i]] = i
    }

    res := make([]int, 0)
    for i := n - 1; i >= 0; i-- {
        num := target - nums[i]
        if j, ok := hash[num]; ok && i != j {
            res = []int{i, j}
        }
    }

    return res
}
