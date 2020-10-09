<?php
class Solution {
    /**
     * @param ListNode $head
     * @return Boolean
     */
    function hasCycle($head) {
        if ($head === null) {
            return false;
        }

        $slow = $head;
        $fast = $head->next;

        while ($fast !== null) {
            if ($slow === $fast) {
                return true;
            }

            $slow = $slow->next;
            $fast = $fast->next;

            if ($fast !== null) {
                $fast = $fast->next;
            }
        }

        return false;
    }
}