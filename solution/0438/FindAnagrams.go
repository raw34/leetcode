package _0438

func findAnagrams(s string, p string) []int {
    res := make([]int, 0)
    len1 := len(p)
    len2 := len(s)
    if len1 > len2 {
        return res
    }
    cnt1 := [26]int{}
    cnt2 := [26]int{}
    for i := 0; i < len1; i++ {
        cnt1[p[i]-'a']++
        cnt2[s[i]-'a']++
    }

    if cnt1 == cnt2 {
        res = append(res, 0)
    }

    for i := len1; i < len2; i++ {
        cnt2[s[i]-'a']++
        cnt2[s[i-len1]-'a']--
        if cnt1 == cnt2 {
            res = append(res, i-len1+1)
        }
    }

    return res
}
