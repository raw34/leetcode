<?php
class ReverseList {
    /**
     * @param ListNode $head
     * @return ListNode
     */
    function reverseList($head) {
        if ($head === null || $head->next === null) {
            return $head;
        }

        $prev = null;

        while ($head) {
            $tmp = $head->next;
            $head->next = $prev;
            $prev = $head;
            $head = $tmp;
        }

        return $prev;
    }
}
