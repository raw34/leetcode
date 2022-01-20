package _0068

import "strings"

func fullJustify(words []string, maxWidth int) []string {
    n := len(words)
    justify := func(i int, words []string) string {
        str := strings.Join(words, " ")
        width := len(str)
        m := len(words)
        j := 1
        for width < maxWidth {
            if m == 1 || i == n-1 {
                // 当前行只有一个单词或当前是最后一行时，字符串尾部补个空格
                str += " "
                width++
            } else {
                // 保证不越界
                if j == m {
                    j = 1
                }
                // 在当前单词前补个空格
                words[j] = " " + words[j]
                str = strings.Join(words, " ")
                width++
                j++
            }
        }

        return str
    }

    res := make([]string, 0)
    for i := 0; i < n; i++ {
        curr := words[i]
        width := len(curr)
        currWords := []string{curr}
        for j := i + 1; j < n; j++ {
            next := words[j]
            width += len(next)
            if width+1 > maxWidth {
                // 追加下一个单词后长度越界时，不选
                break
            }
            // 追加当前单词
            currWords = append(currWords, next)
            width++
            i++
        }
        // 追加当前行处理结果
        res = append(res, justify(i, currWords))
    }

    return res
}
