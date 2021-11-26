package _394

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stack := make([]string, 0)
	for i := 0; i < len(s); i++ {
		curr := string(s[i])
		if curr != "]" {
			stack = append(stack, curr)
		} else {
			prev := ""
			tmpNum := ""
			tmpStr := ""
			for len(stack) > 0 {
				next := stack[len(stack)-1]
				if !(next >= "0" && next <= "9") && (prev >= "0" && prev <= "9") {
					break
				}
				prev = next
				stack = stack[0:len(stack)-1]
				if next == "[" {
					continue
				} else if next >= "0" && next <= "9" {
					tmpNum = next + tmpNum
				} else {
					tmpStr = next + tmpStr
				}
			}
			num, err := strconv.Atoi(tmpNum)
			if err == nil {
				tmpStr = strings.Repeat(tmpStr, num)
			}
			stack = append(stack, tmpStr)
		}
	}

	return strings.Join(stack, "")
}