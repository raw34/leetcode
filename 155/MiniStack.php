<?php
class MinStack {
    private $stack;
    private $stackMin;

    public function __construct()
    {
        $this->stack = [];
        $this->stackMin = [];
    }

    /**
     * @param Integer $x
     * @return NULL
     */
    public function push($x) {
        $this->stack[] = $x;

        if (empty($this->stackMin) || $x <= $this->stackMin[count($this->stackMin) - 1]) {
            $this->stackMin[] = $x;
        }
    }

    /**
     * @return NULL
     */
    public function pop() {
        $x = array_pop($this->stack);

        if ($x === $this->stackMin[count($this->stackMin) - 1]) {
            array_pop($this->stackMin);
        }
    }

    /**
     * @return Integer
     */
    public function top() {
        return $this->stack[count($this->stack) - 1];
    }

    /**
     * @return Integer
     */
    public function getMin() {
        return $this->stackMin[count($this->stackMin) -1];
    }
}

$obj = new MinStack();
$obj->push(-2);
$obj->push(0);
$obj->push(-3);
echo "min = ", $obj->getMin(), "\n";
$obj->pop();
echo "top = ", $obj->top(), "\n";
echo "min = ", $obj->getMin(), "\n";
