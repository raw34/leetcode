<?php
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */

class Codec {
    function __construct() {

    }

    /**
     * @param TreeNode $root
     * @return String
     */
    function serialize($root) {
        $queue = [$root];
        $arr = [];

        while (!empty($queue)) {
            $node = array_shift($queue);

            if ($node) {
                $arr[] = $node->val;
                $queue[] = $node->left;
                $queue[] = $node->right;
            } else {
                $arr[] = null;
            }
        }

        return json_encode($arr);
    }

    /**
     * @param String $data
     * @return TreeNode
     */
    function deserialize($data) {
        $arr = json_decode($data);
        $root = new TreeNode($arr[0]);
        $queue = [$root];

        $i = 1;
        while (!empty($queue)) {
            $node = array_shift($queue);

            if ($arr[$i] !== null) {
                $node->left = new TreeNode($arr[$i]);
                $queue[] = $node->left;
            }
            $i++;

            if ($arr[$i] !== null) {
                $node->right = new TreeNode($arr[$i]);
                $queue[] = $node->right;
            }
            $i++;
        }

        return $root;
    }
}

/**
 * Your Codec object will be instantiated and called as such:
 * $obj = Codec();
 * $data = $obj->serialize($root);
 * $ans = $obj->deserialize($data);
 */