package _0136

func singleNumber(nums []int) int {
    x := 0
    for _, num := range nums {
        x = x ^ num
    }

    return x
}
