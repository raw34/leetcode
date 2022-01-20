package _0043

import "fmt"

func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }

    len1 := len(num1)
    len2 := len(num2)
    // 不能用暴力解法，因为普通计算在数字较大时会有精度问题
    // 利用两数乘积的位数必定小于两数位数之和的规律，构建一个结果数字位数长度的数组
    nums := make([]int, len1+len2)
    for i := len1 - 1; i >= 0; i-- {
        n1 := num1[i] - '0'
        for j := len2 - 1; j >= 0; j-- {
            n2 := num2[j] - '0'
            // 每次计算左右两数一位数字
            sum := nums[i+j+1] + int(n1)*int(n2)
            // 余数赋值给当前位
            nums[i+j+1] = sum % 10
            // 进位数累加到前一位
            nums[i+j] += sum / 10
        }
    }

    // 从头遍历每位数字，拼成计算结果
    res := ""
    for i := 0; i < len(nums); i++ {
        if i == 0 && nums[0] == 0 {
            continue
        }
        res += fmt.Sprintf("%d", nums[i])
    }

    return res
}
