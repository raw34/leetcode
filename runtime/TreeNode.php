<?php

namespace runtime;

/**
 * Class TreeNode
 * @author Randy Chang
 */
class TreeNode
{
    public $val;
    public $left;
    public $right;

    public function __construct($val)
    {
        $this->val = $val;
    }
}
