<?php

use runtime\ListNode;

class Solution {

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2) {
        $cur = new ListNode(0);
        $carry = 0;

        while ($l1 || $l2 || $carry) {
            $x = $l1 ? $l1->val : 0;
            $y = $l2 ? $l2->val : 0;
            $sum = $x + $y + $carry;
            $carry = floor($sum / 10);

//            echo "x = $x, y = $y, sum = $sum, carry = $carry\n";

            $node = new ListNode($sum % 10);

//            echo "val = $node->val\n";

            $cur->next = $node;
            $cur = $cur->next;

            $l1 = $l1 ? $l1->next : $l1;
            $l2 = $l2 ? $l2->next : $l2;
        }

        return $cur->next;
    }
}