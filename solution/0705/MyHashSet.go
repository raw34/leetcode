package _0705

import (
    "container/list"
)

type MyHashSet struct {
    data [1000001]bool
}

func Constructor() MyHashSet {
    return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
    this.data[key] = true
}

func (this *MyHashSet) Remove(key int) {
    this.data[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
    return this.data[key]
}

const base = 997 // 较大的质数

type MyHashSet2 struct {
    data []list.List
}

func (this *MyHashSet2) hash(key int) int {
    return key % base
}

func Constructor2() MyHashSet2 {
    data := make([]list.List, base)
    return MyHashSet2{data}
}

func (this *MyHashSet2) Add(key int) {
    if !this.Contains(key) {
        h := this.hash(key)
        this.data[h].PushBack(key)
    }
}

func (this *MyHashSet2) Remove(key int) {
    h := this.hash(key)
    for e := this.data[h].Front(); e != nil; e = e.Next() {
        if e.Value.(int) == key {
            this.data[h].Remove(e)
        }
    }
}

func (this *MyHashSet2) Contains(key int) bool {
    h := this.hash(key)
    for e := this.data[h].Front(); e != nil; e = e.Next() {
        if e.Value.(int) == key {
            return true
        }
    }
    return false
}
