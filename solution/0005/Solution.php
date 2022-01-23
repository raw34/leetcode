<?php

class Solution
{

    /**
     * @param String $s
     * @return String
     */
    function longestPalindrome($s)
    {
        $len = strlen($s);

        if ($len === 0) {
            return '';
        }

        $res = '';

        for ($i = 0; $i < $len; $i++) {
            $s1 = $this->valid($s, $i, $i);
            $s2 = $this->valid($s, $i, $i + 1);

            $res = strlen($s1) > strlen($res) ? $s1 : $res;
            $res = strlen($s2) > strlen($res) ? $s2 : $res;
        }

        return $res;
    }

    function valid($str, $l, $r)
    {
        while ($l >= 0 && $r < strlen($str) && $str[$l] === $str[$r]) {
            $l--;
            $r++;
        }

        return substr($str, $l + 1, $r - $l - 1);
    }
}