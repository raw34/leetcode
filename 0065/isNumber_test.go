package _0065

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_isNumber(t *testing.T) {
    assert.True(t, isNumber("2"))
    assert.True(t, isNumber("0089"))
    assert.True(t, isNumber("-0.1"))
    assert.True(t, isNumber("+3.14"))
    assert.True(t, isNumber(".1"))
    assert.True(t, isNumber("4."))
    assert.True(t, isNumber("-.9"))
    assert.True(t, isNumber("2e10"))
    assert.True(t, isNumber("-90E3"))
    assert.True(t, isNumber("3e+7"))
    assert.True(t, isNumber("+6e-1"))
    assert.True(t, isNumber("53.5e93"))
    assert.True(t, isNumber("-123.456e789"))
    assert.True(t, isNumber("-90E3"))
    assert.True(t, isNumber("46.e3"))
    assert.True(t, isNumber("+42e+76125"))

    assert.False(t, isNumber(""))
    assert.False(t, isNumber("."))
    assert.False(t, isNumber("abc"))
    assert.False(t, isNumber("1a"))
    assert.False(t, isNumber("1e"))
    assert.False(t, isNumber(".e1"))
    assert.False(t, isNumber("1e."))
    assert.False(t, isNumber("e3"))
    assert.False(t, isNumber("99e2.5"))
    assert.False(t, isNumber("--6"))
    assert.False(t, isNumber("-+3"))
    assert.False(t, isNumber("95a54e53"))
    assert.False(t, isNumber("6+1"))
    assert.False(t, isNumber("+E3"))
    assert.False(t, isNumber("-9e0E3"))
    assert.False(t, isNumber(".-4"))
}
