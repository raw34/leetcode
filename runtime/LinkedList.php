<?php
namespace runtime;

/**
 * Interface LinkedList
 * @author Randy Chang
 */
interface LinkedList
{
    public function init($arr = []);

    public function insert($pos, $item);

    public function delete($pos);

    public function update($pos, $item);

    public function get($pos);

    public function clear();

    public function size();

    public function display();
}

/**
 * Interface AcyclicLinkedList
 * @author Randy Chang
 */
interface AcyclicLinkedList
{
    public function unshift($item);

    public function shift();

    public function push($item);

    public function pop();
}

/**
 * Class ListNode
 * @author Randy Chang
 */
class ListNode
{
    public $val;
    public $prev;
    public $next;

	public function __construct($val)
	{
	    $this->val = $val;
	}
}

/**
 * Class SingleLinkedList
 * @author Randy Chang
 */
class SingleLinkedList implements LinkedList
{
    public $head;

    public function __construct()
    {
        $this->head = new ListNode(null);
    }

    public function init($arr = [])
    {
        foreach ($arr as $k => $v) {
            $this->insert($k, $v);
        }
    }

    public function insert($pos, $item)
    {
        $node = new ListNode($item);

        $curr = $this->get($pos);
        $next = $curr->next;

        $node->next = $next;
        $curr->next = $node;
    }

    public function delete($pos)
    {
        $prev = $this->get($pos - 1);
        $next = $prev->next->next;

        $prev->next = $next;
    }

    public function update($pos, $item)
    {
        $node = new ListNode($item);

        $prev = $this->get($pos - 1);
        $next = $prev->next->next;

        $node->next = $next;
        $prev->next = $node;
    }

    public function get($pos)
    {
        $curr = $this->head;

        $i = 0;
        while ($curr->next && $i <= $pos) {
            $curr = $curr->next;
            $i++;
        }

        return $curr;
    }

    public function reverse()
    {
        $curr = $this->head;
        $prev = null;

        while ($curr) {
            $next = $curr->next;
            $curr->next = $prev;
            $prev = $curr;
            $curr = $next;
        }

        $headNew = new ListNode(null);
        $headNew->next = $prev;

        $this->head = $headNew;
	}

    public function size()
    {
        $length = 0;

        $curr = $this->head;

        while ($curr->next) {
            $curr = $curr->next;
            $length++;
        }

        return $length;
    }

    public function clear()
    {
        $this->head = null;
	}

    public function display()
    {
        $curr = $this->head;

        while ($curr->next) {
            echo "{$curr->next->val} ";
            $curr = $curr->next;
        }

        echo "\n";
    }
}

/**
 * Class DoubleLinkedList
 * @author Randy Chang
 */
class DoubleLinkedList extends SingleLinkedList
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


/**
 * Class CircularLinkedList
 * @author Randy Chang
 */
class CircularLinkedList
{
}

$list = new SingleLinkedList();
//$list = new DoubleLinkedList();

$list->init([1, 3, 4, 2]);
echo "init 1, 3, 4, 2\n";
$list->display();

$list->insert(3, 32);
echo "insert 32 after pos 3\n";
$list->display();

$list->delete(3);
echo "delete pos 3\n";
$list->display();

$list->update( 3, 25);
echo "update pos 3 to 25\n";
$list->display();

echo "get pos 3\n";
echo $list->get(3)->val, "\n";

$list->reverse();
echo "reverse\n";
$list->display();

