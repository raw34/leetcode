<?php

namespace runtime;

use runtime\ListNode;
use runtime\SingleLinkedList;

/**
 * Class DoubleLinkedList
 * @author Randy Chang
 */
class DLinkedList extends SingleLinkedList
{
    public $tail;

    public function insert($pos, $item)
    {
        $node = new ListNode($item);

        $curr = $this->get($pos);
        $next = $curr->next;

        $node->prev = $curr;
        $node->next = $next;
        $curr->next = $node;

        if ($next !== null) {
            $next->prev = $node;
        } else {
            $this->tail = $node;
        }
    }

    public function delete($pos)
    {
        $prev = $this->get($pos - 1);
        $next = $prev->next->next;

        $prev->next = $next;

        if ($next !== null) {
            $next->prev = $prev;
        } else {
            $this->tail = $prev;
        }
    }

    public function update($pos, $item)
    {
        $node = new ListNode($item);

        $prev = $this->get($pos - 1);
        $next = $prev->next->next;

        $node->prev = $prev;
        $node->next = $next;
        $prev->next = $node;

        if ($next !== null) {
            $next->prev = $node;
        } else {
            $this->tail = $node;
        }
    }

    public function displayReverse()
    {
        $curr = $this->tail;

        while ($curr->prev) {
            echo $curr->prev->val, ' ';
            $curr = $curr->prev;
        }

        echo "\n";
    }
}