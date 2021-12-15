package _1028

import (
    "fmt"
    "testing"
)

func Test_recoverFromPreorder(t *testing.T) {
    str := "1-2--3--4-5--6--7"
    res := recoverFromPreorder(str)

    fmt.Println("res", res)
}
