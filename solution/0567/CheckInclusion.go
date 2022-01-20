package _0567

func checkInclusion(s1 string, s2 string) bool {
    n1 := len(s1)
    n2 := len(s2)
    if n1 > n2 {
        return false
    }

    cnt1 := [26]int{}
    cnt2 := [26]int{}
    for i := 0; i < n1; i++ {
        cnt1[s1[i]-'a']++
        cnt2[s2[i]-'a']++
    }

    if cnt1 == cnt2 {
        return true
    }

    for i := n1; i < n2; i++ {
        cnt2[s2[i]-'a']++
        cnt2[s2[i-n1]-'a']--
        if cnt1 == cnt2 {
            return true
        }
    }

    return false
}
