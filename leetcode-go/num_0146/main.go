package main

import (
	"fmt"
)

/**

146. LRU 缓存

请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value；
如果不存在，则向缓存中插入该组 key-value。如果插入操作导致关键字数量超过 capacity，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。


示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4


提示：

1 <= capacity <= 3000
0 <= key <= 10000
0 <= value <= 105
最多调用 2 * 105 次 get 和 put

*/

type LRUCache struct {
	capacity   int
	linkedList *DoublyLinkedList
	cache      map[int]*DoublyListNode
}

type DoublyListNode struct {
	key, val   int
	prev, next *DoublyListNode
}

type DoublyLinkedList struct {
	head, tail *DoublyListNode
}

func newDoublyLinkedList() *DoublyLinkedList {
	head := &DoublyListNode{}
	tail := &DoublyListNode{}
	head.next = tail
	tail.prev = head
	return &DoublyLinkedList{head, tail}
}

func (l *DoublyLinkedList) moveToHead(node *DoublyListNode) {
	if node == nil {
		return
	}
	// 断开当前结点
	node.next.prev = node.prev
	node.prev.next = node.next
	// 节点移动到表头：最新访问
	l.addNode(node)
}

// addNode 添加的新结点始终在头结点后面
func (l *DoublyLinkedList) addNode(node *DoublyListNode) {
	if node == nil {
		return
	}
	node.next = l.head.next
	node.prev = l.head
	node.next.prev = node
	l.head.next = node
}

func (l *DoublyLinkedList) removeLastNode() {
	lastNode := l.tail.prev
	lastNode.prev.next = l.tail
	l.tail.prev = lastNode.prev
}

func Constructor(capacity int) LRUCache {
	cache := make(map[int]*DoublyListNode)
	linkedList := newDoublyLinkedList()
	return LRUCache{capacity, linkedList, cache}
}

func (c *LRUCache) Get(key int) int {
	values := c.cache
	node, ok := values[key]
	if !ok {
		return -1
	}

	c.linkedList.moveToHead(node)
	return node.val
}

func (c *LRUCache) Put(key int, value int) {
	cache := c.cache
	node, ok := cache[key]
	// 更新最近访问节点
	if ok {
		// 更新结点值
		node.val = value
		// 移动到表头
		c.linkedList.moveToHead(node)
		return
	}
	// 新增节点，判断是否需要淘汰最久未使用节点
	newNode := &DoublyListNode{key, value, nil, nil}
	cache[key] = newNode
	c.linkedList.addNode(newNode)
	if len(cache) <= c.capacity {
		return
	}
	// 移除最久未使用节点
	lastNode := c.linkedList.tail.prev
	c.linkedList.removeLastNode()
	delete(cache, lastNode.key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)

	fmt.Println(lru.Get(1))
	lru.Put(3, 3)

	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))
	fmt.Println(lru.cache)
	fmt.Println(lru.linkedList.head.next.val)
}
