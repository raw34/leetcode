package _0012

func intToRoman(num int) string {
    if num == 0 {
        return "0"
    }

    nums1 := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    nums2 := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    res := ""
    n := len(nums1)
    for i := 0; i < n; {
        if num == 0 {
            break
        }
        n1 := nums1[i]
        n2 := nums2[i]
        if num >= n1 {
            res += n2
            num -= n1
        } else {
            i++
        }
    }

    return res
}
