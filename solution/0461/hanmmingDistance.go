package _0461

func hammingDistance(x int, y int) int {
    res := 0
    z := x ^ y
    for z != 0 {
        z = z & (z - 1)
        res++
    }

    return res
}
