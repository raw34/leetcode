package runtime

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestLinkedList_InOrder1(t *testing.T) {
    //["LinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
    //[[],[1],[3],[1,2],[1],[1],[1]]
    list := NewLinkedList()
    list.AddAtHead(1)
    list.AddAtTail(3)
    list.AddAtIndex(1, 2)
    assert.Equal(t, 2, list.Get(1))
    list.DeleteAtIndex(1)
    assert.Equal(t, 3, list.Get(1))
}

func TestLinkedList_InOrder2(t *testing.T) {
    //["LinkedList", "addAtHead", "deleteAtIndex"]
    //[[], [1], [0]]
    list := NewLinkedList()
    list.AddAtHead(1)
    list.DeleteAtIndex(0)
    assert.Equal(t, -1, list.Get(0))
}

func TestLinkedList_InOrder3(t *testing.T) {
    //["LinkedList","addAtIndex","addAtIndex","addAtIndex","get"]
    //[[],[0,10],[0,20],[1,30],[0]]
    list := NewLinkedList()
    list.AddAtIndex(0, 10)
    list.AddAtIndex(0, 20)
    list.AddAtIndex(1, 30)
    assert.Equal(t, 20, list.Get(0))
}

func TestLinkedList_InOrder4(t *testing.T) {
    //["LinkedList","addAtTail","get"]
    //[[],[1],[0]]
    list := NewLinkedList()
    list.AddAtTail(1)
    assert.Equal(t, 1, list.Get(0))
}

func TestLinkedList_InOrder5(t *testing.T) {
    //["LinkedList","addAtHead","deleteAtIndex","addAtHead","addAtHead","addAtHead","addAtHead","addAtHead","addAtTail","get","deleteAtIndex","deleteAtIndex"]
    //[[],[2],[1],[2],[7],[3],[2],[5],[5],[5],[6],[4]]
    list := NewLinkedList()
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

func TestLinkedList_InOrder6(t *testing.T) {
    //["LinkedList","addAtHead","get","addAtHead","addAtHead","deleteAtIndex","addAtHead","get","get","get","addAtHead","deleteAtIndex"]
    //[[],[4],[1],[1],[5],[3],[7],[3],[3],[3],[1],[4]]
    list := NewLinkedList()
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

func TestLinkedList_InOrder7(t *testing.T) {
    //["LinkedList","addAtHead","addAtIndex","get"]
    //[[],[2],[0,1],[1]]
    list := NewLinkedList()
    list.AddAtHead(2)
    list.AddAtIndex(0, 1)
    assert.Equal(t, 2, list.Get(1))
}

func TestLinkedList_InOrder8(t *testing.T) {
    //["LinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get","get","deleteAtIndex","deleteAtIndex","get","deleteAtIndex","get"]
    //[[],[1],[3],[1,2],[1],[1],[1],[3],[3],[0],[0],[0],[0]]
    list := NewLinkedList()
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

func TestLinkedList_InOrder9(t *testing.T) {
    //["LinkedList","addAtHead","addAtTail","addAtIndex"]
    //[[],[1],[3],[3,2]]
    list := NewLinkedList()
    list.AddAtHead(1)
    list.AddAtTail(3)
    list.AddAtIndex(3, 2)
    assert.Equal(t, -1, list.Get(3))
}

func TestLinkedList_InOrder10(t *testing.T) {
    //["LinkedList","addAtIndex","get"]
    //[[],[1,0],[0]]
    list := NewLinkedList()
    list.AddAtIndex(1, 0)
    assert.Equal(t, -1, list.Get(0))
}

func TestLinkedList_InOrder11(t *testing.T) {
    //["LinkedList","addAtHead","deleteAtIndex","addAtHead","addAtHead","addAtHead","addAtHead","addAtHead","addAtTail","get","deleteAtIndex","deleteAtIndex"]
    //[[],[2],[1],[2],[7],[3],[2],[5],[5],[5],[6],[4]]
    list := NewLinkedList()
    list.AddAtHead(2)
    list.Display()
    list.DeleteAtIndex(1)
    list.Display()
    list.AddAtHead(2)
    list.AddAtHead(7)
    list.AddAtHead(3)
    list.AddAtHead(2)
    list.AddAtHead(5)
    list.AddAtTail(5)
    list.Display()
    assert.Equal(t, 2, list.Get(5))
    list.DeleteAtIndex(6)
    list.DeleteAtIndex(4)
}
