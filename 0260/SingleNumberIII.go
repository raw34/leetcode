package _0160

func singleNumberIII(nums []int) []int {
    xorSum := 0
    for _, num := range nums {
        xorSum ^= num
    }
    lsb := xorSum & -xorSum
    type1, type2 := 0, 0
    for _, num := range nums {
        if num&lsb > 0 {
            type1 ^= num
        } else {
            type2 ^= num
        }
    }

    return []int{type1, type2}
}
