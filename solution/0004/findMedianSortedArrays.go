package _0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    n := len(nums1) + len(nums2)
    mid := float64(n) / 2
    midInt := int(mid)
    odd := float64(midInt) != mid
    median := float64(0)
    i := 0
    for len(nums1) > 0 || len(nums2) > 0 {
        if i > midInt {
            break
        }
        var num int
        if len(nums1) > 0 && len(nums2) > 0 {
            num = min(nums1[0], nums2[0])
            if nums1[0] == num {
                nums1 = nums1[1:]
            } else {
                nums2 = nums2[1:]
            }
        } else if len(nums1) > 0 {
            num = nums1[0]
            nums1 = nums1[1:]
        } else if len(nums2) > 0 {
            num = nums2[0]
            nums2 = nums2[1:]
        }
        if odd && i == midInt {
            median += float64(num)
        }
        if !odd && (i == midInt-1 || i == midInt) {
            median += float64(num)
        }
        i++
    }

    if !odd {
        median = float64(median) / 2
    }

    return median
}
