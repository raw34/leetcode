package _0074

func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    left := 0
    right := m*n - 1

    for left <= right {
        mid := (left + right) / 2
        x := matrix[mid/n][mid%n]
        if x < target {
            left = mid + 1
        } else if x > target {
            right = mid - 1
        } else {
            return true
        }
    }

    return false
}
