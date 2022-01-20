package _0022

func generateParenthesis(n int) []string {
    res := make([]string, 0)
    var dfs func(path string, left, right int)
    dfs = func(path string, left, right int) {
        if left == n && right == n {
            res = append(res, path)
            return
        }
        if left < right {
            return
        }
        if left < n {
            dfs(path+"(", left+1, right)
        }
        if right < n {
            dfs(path+")", left, right+1)
        }
    }
    dfs("", 0, 0)

    return res
}
