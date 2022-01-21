package _0349

func intersection(nums1 []int, nums2 []int) []int {
    res := make([]int, 0)
    hash := map[int]int{}
    for i := 0; i < len(nums1); i++ {
        num := nums1[i]
        if hash[num] == 0 {
            hash[num]++
        }
    }
    for j := 0; j < len(nums2); j++ {
        num := nums2[j]
        if hash[num] == 1 {
            res = append(res, num)
            hash[num]++
        }
    }

    return res
}
