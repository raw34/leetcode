package _0370

func getModifiedArray(length int, updates [][]int) []int {
    res := make([]int, length)
    for _, update := range updates {
        res[update[0]] += update[2]
        if update[1] < length-1 {
            res[update[1]+1] -= update[2]
        }
    }
    for i := 1; i < length; i++ {
        res[i] += res[i-1]
    }

    return res
}
