package _0017

func letterCombinations(digits string) []string {
    res := make([]string, 0)
    if digits == "" {
        return res
    }

    hash := map[byte]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    var backtrack func(list []byte, start int)
    backtrack = func(list []byte, start int) {
        if len(list) == len(digits) {
            res = append(res, string(list))
            return
        }
        letters := hash[digits[start]]
        for i := 0; i < len(letters); i++ {
            list = append(list, letters[i])
            backtrack(list, start+1)
            list = list[0 : len(list)-1]
        }
    }
    backtrack([]byte{}, 0)

    return res
}
