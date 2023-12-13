/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start
type LRUCache struct {
    size int
    capacity int
    // 利用map快速找到结点
    cache map[int]*DLinkedNode
    // 核心需要两个头尾结点
    head, tail *DLinkedNode
}

type DLinkedNode struct {
    key, val int
    prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
    head := &DLinkedNode{}
    tail := &DLinkedNode{}
    head.next = tail
    tail.prev = head
    return LRUCache{
        capacity: capacity,
        cache: make(map[int]*DLinkedNode),
        head: head,
        tail: tail,
    }
}


func (this *LRUCache) Get(key int) int {
    if _, ok := this.cache[key]; !ok {
        return -1
    }
    node := this.cache[key]
    this.moveToHead(node)
    return node.val
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
    node.next.prev = node.prev
    node.prev.next = node.next
}

func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.cache[key]; !ok {
        // 不存在
        node := &DLinkedNode{
            key: key,
            val: value,
        }
        this.cache[key] = node
        this.addToHead(node)
        this.size++
        if this.size > this.capacity {
             removeNode := this.removeTail()
             delete(this.cache, removeNode.key)
             this.size--
        }
    }else {
        // 存在
        node := this.cache[key]
        node.val = value
        this.moveToHead(node)
    }
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
    this.removeNode(node)
    this.addToHead(node)
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
    node.next = this.head.next
    this.head.next.prev = node
    this.head.next = node
    node.prev = this.head
}

func (this *LRUCache) removeTail() (*DLinkedNode) {
    node := this.tail.prev
    this.removeNode(node)
    return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

