package _0704

func search(nums []int, target int) int {
    var dfs func(left, right int) int
    dfs = func(left, right int) int {
        if left > right {
            return -1
        }
        mid := (left + right) / 2
        curr := nums[mid]
        if curr == target {
            return mid
        } else if curr > target {
            return dfs(left, mid-1)
        } else {
            return dfs(mid+1, right)
        }
    }

    return dfs(0, len(nums)-1)
}

func search2(nums []int, target int) int {
    left := 0
    right := len(nums) - 1

    for left <= right {
        mid := (left + right) / 2
        curr := nums[mid]
        if curr == target {
            return mid
        } else if curr > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return -1
}
