package _0088

func merge(nums1 []int, m int, nums2 []int, n int) {
    nums3 := make([]int, m+n)
    i, j, k := 0, 0, 0
    for i < m || j < n {
        num := 0
        if i == m {
            num = nums2[j]
            j++
        } else if j == n {
            num = nums1[i]
            i++
        } else if nums1[i] < nums2[j] {
            num = nums1[i]
            i++
        } else {
            num = nums2[j]
            j++
        }
        nums3[k] = num
        k++
    }
    for i := 0; i < m+n; i++ {
        nums1[i] = nums3[i]
    }
}
