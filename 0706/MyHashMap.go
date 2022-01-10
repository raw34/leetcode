package _0706

import "container/list"

type MyHashMap struct {
    data []int
}

func Constructor() MyHashMap {
    data := make([]int, 1000001)
    for i := 0; i < len(data); i++ {
        data[i] = -1
    }

    return MyHashMap{data}
}

func (this *MyHashMap) Put(key int, value int) {
    this.data[key] = value
}

func (this *MyHashMap) Get(key int) int {
    return this.data[key]
}

func (this *MyHashMap) Remove(key int) {
    this.data[key] = -1
}

const base = 997

type entry struct {
    key, value int
}

type MyHashMap2 struct {
    data []list.List
}

func Constructor2() MyHashMap2 {
    data := make([]list.List, base)
    return MyHashMap2{data}
}

func (this *MyHashMap2) hash(key int) int {
    return key % base
}

func (this *MyHashMap2) Put(key int, value int) {
    h := this.hash(key)
    for e := this.data[h].Front(); e != nil; e = e.Next() {
        if et := e.Value.(entry); et.key == key {
            e.Value = entry{key, value}
            return
        }
    }
    this.data[h].PushBack(entry{key, value})
}

func (this *MyHashMap2) Get(key int) int {
    h := this.hash(key)
    for e := this.data[h].Front(); e != nil; e = e.Next() {
        if et := e.Value.(entry); et.key == key {
            return et.value
        }
    }
    return -1
}

func (this *MyHashMap2) Remove(key int) {
    h := this.hash(key)
    for e := this.data[h].Front(); e != nil; e = e.Next() {
        if e.Value.(entry).key == key {
            this.data[h].Remove(e)
        }
    }
}
