package _0015

import "sort"

func threeSum(nums []int) [][]int {
    // 预处理，先排序，以便后续计算方便
    sort.Ints(nums)
    res := make([][]int, 0)
    n := len(nums)
    // 从头遍历每个数字
    for i := 0; i < n; i++ {
        // 由于数组已为升序，当前数字大于0时，不可能找到符合条件的数字，直接退出
        if nums[i] > 0 {
            break
        }
        // 当前一个数字已经找过的时，排重
        if i > 0 && nums[i-1] == nums[i] {
            continue
        }

        // 利用双指针同时寻找另两个符合条件的数字
        l := i + 1
        r := n - 1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if sum > 0 {
                // 当三数和大于0时，说明右侧数大了，右指针左移
                r--
            } else if sum < 0 {
                // 当三数和小于0时，说明左侧数小了，左指针右移
                l++
            } else {
                // 符合条件时，结果集追加新子集
                res = append(res, []int{nums[i], nums[l], nums[r]})
                // 左右指针收拢，对相同数字做排重处理
                l++
                r--
                for l < r && nums[l-1] == nums[l] {
                    l++
                }
                for l < r && nums[r+1] == nums[r] {
                    r--
                }
            }
        }
    }

    return res
}
