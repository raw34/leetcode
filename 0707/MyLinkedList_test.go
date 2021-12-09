package _0707

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestMyLinkedList_InOrder1(t *testing.T) {
    //["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
    //[[],[1],[3],[1,2],[1],[1],[1]]
    list := Constructor()
    list.AddAtHead(1)
    list.AddAtTail(3)
    list.AddAtIndex(1, 2)
    assert.Equal(t, 2, list.Get(1))
    list.DeleteAtIndex(1)
    assert.Equal(t, 3, list.Get(1))
}

func TestMyLinkedList_InOrder2(t *testing.T) {
    //["MyLinkedList", "addAtHead", "deleteAtIndex"]
    //[[], [1], [0]]
    list := Constructor()
    list.AddAtHead(1)
    list.DeleteAtIndex(0)
    assert.Nil(t, list.Head)
}

func TestMyLinkedList_InOrder3(t *testing.T) {
    //["MyLinkedList","addAtIndex","addAtIndex","addAtIndex","get"]
    //[[],[0,10],[0,20],[1,30],[0]]
    list := Constructor()
    list.AddAtIndex(0, 10)
    list.AddAtIndex(0, 20)
    list.AddAtIndex(1, 30)
    assert.Equal(t, 20, list.Get(0))
}

func TestMyLinkedList_InOrder4(t *testing.T) {
    //["MyLinkedList","addAtTail","get"]
    //[[],[1],[0]]
    list := Constructor()
    list.AddAtTail(1)
    assert.Equal(t, 1, list.Get(0))
}

func TestMyLinkedList_InOrder5(t *testing.T) {
    //["MyLinkedList","addAtHead","deleteAtIndex","addAtHead","addAtHead","addAtHead","addAtHead","addAtHead","addAtTail","get","deleteAtIndex","deleteAtIndex"]
    //[[],[2],[1],[2],[7],[3],[2],[5],[5],[5],[6],[4]]
    list := Constructor()
    list.AddAtHead(2)
    list.DeleteAtIndex(1)
    list.AddAtHead(2)
    list.AddAtHead(7)
    list.AddAtHead(3)
    list.AddAtHead(2)
    list.AddAtHead(5)
    list.AddAtTail(5)
    list.Get(5)
    list.DeleteAtIndex(6)
    list.DeleteAtIndex(4)
}

func TestMyLinkedList_InOrder6(t *testing.T) {
    //["MyLinkedList","addAtHead","get","addAtHead","addAtHead","deleteAtIndex","addAtHead","get","get","get","addAtHead","deleteAtIndex"]
    //[[],[4],[1],[1],[5],[3],[7],[3],[3],[3],[1],[4]]
    list := Constructor()
    list.AddAtHead(4)
    assert.Equal(t, -1, list.Get(1))
    list.AddAtHead(1)
    list.AddAtHead(5)
    list.DeleteAtIndex(3)
    list.AddAtHead(7)
    assert.Equal(t, 4, list.Get(3))
    assert.Equal(t, 4, list.Get(3))
    assert.Equal(t, 4, list.Get(3))
    list.AddAtHead(1)
    list.DeleteAtIndex(4)
}

func TestMyLinkedList_InOrder7(t *testing.T) {
    //["MyLinkedList","addAtHead","addAtIndex","get"]
    //[[],[2],[0,1],[1]]
    list := Constructor()
    list.AddAtHead(2)
    list.AddAtIndex(0, 1)
    assert.Equal(t, 2, list.Get(1))
}

func TestMyLinkedList_InOrder8(t *testing.T) {
    //["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get","get","deleteAtIndex","deleteAtIndex","get","deleteAtIndex","get"]
    //[[],[1],[3],[1,2],[1],[1],[1],[3],[3],[0],[0],[0],[0]]
    list := Constructor()
    list.AddAtHead(1)
    list.AddAtTail(3)
    list.AddAtIndex(1, 2)
    assert.Equal(t, 2, list.Get(1))
    list.DeleteAtIndex(1)
    assert.Equal(t, 3, list.Get(1))
    assert.Equal(t, -1, list.Get(3))
    list.DeleteAtIndex(3)
    list.DeleteAtIndex(0)
    assert.Equal(t, 3, list.Get(0))
    list.DeleteAtIndex(0)
    assert.Equal(t, -1, list.Get(0))
}

func TestMyLinkedList_InOrder9(t *testing.T) {
    //["MyLinkedList","addAtHead","addAtTail","addAtIndex"]
    //[[],[1],[3],[3,2]]
    list := Constructor()
    list.AddAtHead(1)
    list.AddAtTail(3)
    list.AddAtIndex(3, 2)
    assert.NotNil(t, list.Head)
}

func TestMyLinkedList_InOrder10(t *testing.T) {
    //["MyLinkedList","addAtIndex","get"]
    //[[],[1,0],[0]]
    list := Constructor()
    list.AddAtIndex(1, 0)
    assert.Equal(t, -1, list.Get(0))
}
