package _0217

func containsDuplicate(nums []int) bool {
    hash := map[int]bool{}
    for _, num := range nums {
        if hash[num] {
            return true
        }
        hash[num] = true
    }

    return false
}
