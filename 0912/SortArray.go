package _0912

func sortArray(nums []int) []int {
    swap := func(i, j int) {
        nums[i], nums[j] = nums[j], nums[i]
    }

    var down func(n, i int)
    down = func(n, i int) {
        l := 2*i + 1
        r := 2*i + 2
        max := i

        if l < n && nums[l] > nums[max] {
            max = l
        }

        if r < n && nums[r] > nums[max] {
            max = r
        }

        if max != i {
            swap(i, max)
            down(n, max)
        }
    }

    n := len(nums)
    for i := n / 2; i >= 0; i-- {
        down(n, i)
    }
    for i := n - 1; i >= 0; i-- {
        swap(i, 0)
        down(i, 0)
    }

    return nums
}
