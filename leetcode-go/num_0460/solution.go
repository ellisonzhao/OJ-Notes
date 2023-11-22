package main

import (
	"fmt"
)

/**

460. LFU 缓存

请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。

实现 LFUCache 类：

LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1。
void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。
当缓存达到其容量 capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，
应该去除 最久未使用 的键。为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器。使用计数最小的键是最久未使用的键。

当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。

函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。



示例：

输入：
["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
输出：
[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

解释：
// cnt(x) = 键 x 的使用计数
// cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
LFUCache lfu = new LFUCache(2);
lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
lfu.get(1);      // 返回 1
                 // cache=[1,2], cnt(2)=1, cnt(1)=2
lfu.put(3, 3);   // 去除键 2，因为 cnt(2)=1，使用计数最小
                 // cache=[3,1], cnt(3)=1, cnt(1)=2
lfu.get(2);      // 返回 -1（未找到）
lfu.get(3);      // 返回 3
                 // cache=[3,1], cnt(3)=2, cnt(1)=2
lfu.put(4, 4);   // 去除键 1，1 和 3 的 cnt 相同，但 1 最久未使用
                 // cache=[4,3], cnt(4)=1, cnt(3)=2
lfu.get(1);      // 返回 -1（未找到）
lfu.get(3);      // 返回 3
                 // cache=[3,4], cnt(4)=1, cnt(3)=3
lfu.get(4);      // 返回 4
                 // cache=[3,4], cnt(4)=2, cnt(3)=3


提示：

1 <= capacity <= 104
0 <= key <= 105
0 <= value <= 109
最多调用 2 * 105 次 get 和 put 方法

*/

type LFUCache struct {
	capacity, minFreq int
	cache             map[int]*Node
	frequencies       map[int]*DoublyLinkedList
}

type Node struct {
	key, value, freq int
	prev, next       *Node
}

type DoublyLinkedList struct {
	head, tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &DoublyLinkedList{head, tail}
}

func (l *DoublyLinkedList) addToHead(node *Node) {
	if node == nil {
		return
	}
	firstNode := l.head.next
	l.head.next = node
	node.prev = l.head
	node.next = firstNode
	firstNode.prev = node
}

func (l *DoublyLinkedList) delete(node *Node) {
	if node == nil {
		return
	}
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (l *DoublyLinkedList) isEmpty() bool {
	return l.head.next == l.tail
}

func (l *DoublyLinkedList) removeFromTail() *Node {
	if l.isEmpty() {
		return nil
	}
	node := l.tail.prev
	l.delete(node)
	return node
}

func Constructor(capacity int) LFUCache {
	cache := make(map[int]*Node)
	freq := make(map[int]*DoublyLinkedList)
	return LFUCache{capacity, 0, cache, freq}
}

func (c *LFUCache) Get(key int) int {
	node, ok := c.cache[key]
	if !ok {
		return -1
	}

	c.update(node)

	return node.value
}

// 插入新结点或访问已有结点时，访问频次更新
func (c *LFUCache) update(node *Node) {
	// 当前访问频次对应的链表需要删除该结点
	if oldList, ok := c.frequencies[node.freq]; ok {
		oldList.delete(node)
		// 更新最小访问频次
		if c.minFreq == node.freq && oldList.isEmpty() {
			c.minFreq++
		}
	}
	node.freq++
	// 访问频次 +1 对应链表需要新增/更新
	newList, ok := c.frequencies[node.freq]
	if !ok {
		newList = NewDoublyLinkedList()
	}
	newList.addToHead(node)
	c.frequencies[node.freq] = newList
}

func (c *LFUCache) Put(key int, value int) {
	node, ok := c.cache[key]
	// 节点已存在只需要更新值和访问频次链表
	if ok {
		node.value = value
		c.update(node)
		return
	}
	// 节点不存在可能需要新建链表
	if len(c.cache) == c.capacity {
		// 移除最不经常访问的节点
		if list, ok := c.frequencies[c.minFreq]; ok {
			oldNode := list.removeFromTail()
			delete(c.cache, oldNode.key)
		}
	}
	newNode := &Node{key, value, 1, nil, nil}
	newList, ok := c.frequencies[newNode.freq]
	if !ok {
		newList = NewDoublyLinkedList()
		c.frequencies[newNode.freq] = newList
	}
	newList.addToHead(newNode)
	c.cache[key] = newNode
	c.minFreq = newNode.freq
}

func main() {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	fmt.Println(lfu.Get(1))

	lfu.Put(3, 3)
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(3))

	lfu.Put(4, 4)
	fmt.Println(lfu.Get(1))
	fmt.Println(lfu.Get(3))
	fmt.Println(lfu.Get(4))
}
