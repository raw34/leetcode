package _0013

func romanToInt(s string) int {
    hash := map[string]int{
        "I":  1,
        "IV": 4,
        "V":  5,
        "IX": 9,
        "X":  10,
        "XL": 40,
        "L":  50,
        "XC": 90,
        "C":  100,
        "CD": 400,
        "D":  500,
        "CM": 900,
        "M":  1000,
    }
    res := 0
    n := len(s)
    for i := 0; i < n; i++ {
        j := i + 2
        val := hash[string(s[i])]
        if i < n-1 && hash[s[i:j]] > 0 {
            val = hash[s[i:j]]
            i++
        }

        res += val
    }

    return res
}
