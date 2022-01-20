package _0224

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_calculate_Case1(t *testing.T) {
    res := calculate("1 + 1")
    assert.Equal(t, 2, res)
}

func Test_calculate_Case2(t *testing.T) {
    res := calculate(" 2-1 + 2 ")
    assert.Equal(t, 3, res)
}

func Test_calculate_Case3(t *testing.T) {
    res := calculate("(1+(4+5+2)-3)+(6+8)")
    assert.Equal(t, 23, res)
}

func Test_calculate_Case4(t *testing.T) {
    res := calculate("1-(-2)")
    assert.Equal(t, 3, res)
}

func Test_calculate_Case5(t *testing.T) {
    res := calculate("- (3 - (- (4 + 5) ) )")
    assert.Equal(t, -12, res)
}

func Test_calculate_Case6(t *testing.T) {
    res := calculate("2147483647")
    assert.Equal(t, 2147483647, res)
}
