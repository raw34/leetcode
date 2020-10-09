<?php
class Solution {

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function mergeTwoLists($l1, $l2) {
        if ($l1 === null) {
            return $l2;
        }

        if ($l2 === null) {
            return $l1;
        }

        $head = new ListNode();
        $cur = $head;

        while ($l1 && $l2) {
            if ($l1->val < $l2->val) {
                $cur->next = $l1;
                $l1 = $l1->next;
            } else {
                $cur->next = $l2;
                $l2 = $l2->next;
            }

            $cur = $cur->next;
        }

        $cur->next = $l1 === null ? $l2 : $l1;

        return $head->next;
    }
}
