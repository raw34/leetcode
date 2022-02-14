<?php

namespace runtime;

use PHPUnit\Framework\TestCase;

class BinarySearchTreeTest extends TestCase
{
    public function test01()
    {
        $arr = range(1, 10);
        shuffle($arr);
        $bst = new BinarySearchTree();
        $root = $bst->build($arr);
        $bst->display($root);

        self::assertEquals(1, $bst->getMin($root)->val);
        self::assertEquals(10, $bst->getMax($root)->val);
    }
}
