package _1055

func shortestWay(source string, target string) int {
    res := 0
    j := 0
    for j < len(target) {
        i := 0
        k := j
        // 从头遍历source，利用贪心算法求解
        for i < len(source) && j < len(target) {
            if source[i] == target[j] {
                // source和target当前字符相同时，source指针同时前进
                i++
                j++
            } else {
                // source和target当前字符不同时，target指针同时前进
                i++
            }
        }
        // 遍历完source，j没有增加，说明有不存在的字符，直接返回
        if k == j {
            return -1
        }
        res++
    }

    return res
}
