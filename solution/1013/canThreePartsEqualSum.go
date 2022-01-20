package _1013

func canThreePartsEqualSum(arr []int) bool {
    n := len(arr)
    // 预处理，计算前缀和
    preSum := make([]int, n+1)
    for i := 1; i < n+1; i++ {
        preSum[i] += preSum[i-1] + arr[i-1]
    }
    if preSum[n]%3 != 0 {
        return false
    }

    for i := 1; i < n; i++ {
        left := preSum[i] - preSum[0]
        for j := i + 1; j < n; j++ {
            mid := preSum[j] - preSum[i]
            right := preSum[n] - preSum[j]
            if left == mid && mid == right {
                return true
            }
        }
    }

    return false
}

func canThreePartsEqualSum2(arr []int) bool {
    n := len(arr)
    preSum := make([]int, n+1)
    for i := 1; i < n+1; i++ {
        num := arr[i-1]
        preSum[i] += preSum[i-1] + num
    }
    if preSum[n]%3 != 0 {
        return false
    }

    l, r := 1, n-1
    for l < r {
        mid := preSum[n] / 3
        left := preSum[l] - preSum[0]
        right := preSum[n] - preSum[r]
        if left == mid && right == mid {
            return true
        }
        if left != mid {
            l++
        }
        if right != mid {
            r--
        }
    }

    return false
}
