package _0451

import "bytes"

func frequencySort(s string) string {
    maxCount := 0
    counter := map[byte]int{}
    for i := range s {
        maxCount++
        counter[s[i]]++
    }

    buckets := make([][]byte, maxCount+1)
    for c, cnt := range counter {
        buckets[cnt] = append(buckets[cnt], c)
    }

    res := make([]byte, 0)
    for i := maxCount; i > 0; i-- {
        for _, c := range buckets[i] {
            res = append(res, bytes.Repeat([]byte{c}, i)...)
        }
    }

    return string(res)
}
