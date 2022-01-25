package _1108

func defangIPaddr(address string) string {
    cnt := 0
    for i := 0; i < len(address); i++ {
        if cnt == 4 {
            break
        }
        curr := address[i]
        if curr == '.' {
            address = address[:i] + "[.]" + address[i+1:]
            i += 2
            cnt++
        }
    }

    return address
}
