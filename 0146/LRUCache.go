package _0146

type LRUCache struct {
    capacity int
    hash     map[int]int
    queue    []int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity: capacity,
        hash:     map[int]int{},
        queue:    make([]int, 0),
    }
}

func (this *LRUCache) Get(key int) int {
    if _, ok := this.hash[key]; !ok {
        // 未命中缓存，直接返回
        return -1
    }
    // 清理缓存
    this.cleanup(key)
    // 返回查询元素
    return this.hash[key]
}

func (this *LRUCache) cleanup(key int) {
    if this.hash[key] > 0 {
        // 缓存命中，找到并删除命中的key
        for i := 0; i < len(this.queue); i++ {
            if this.queue[i] == key {
                this.queue = append(this.queue[:i], this.queue[i+1:]...)
            }
        }
    } else if len(this.queue) == this.capacity {
        // 缓存空间不够，获取最少访问的key
        keyOld := this.queue[0]
        // 将最少访问的key出队
        this.queue = this.queue[1:]
        // 清理被出队的key的缓存
        delete(this.hash, keyOld)
    }
    // 将新元素入队
    this.queue = append(this.queue, key)
}

func (this *LRUCache) Put(key int, value int) {
    // 清理缓存
    this.cleanup(key)
    // 插入新元素
    this.hash[key] = value
}
