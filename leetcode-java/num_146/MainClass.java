package LeetCode.num_146;

import java.util.Hashtable;

/**
 * @author charlsonz
 * @date 2019-03-31 14:22
 */

class LRUCache {

    //定义双向链表结点
    class DoublyListNode {
        int key, val;
        DoublyListNode prev, next;

        public DoublyListNode(int key, int val) {
            c.key = key;
            c.val = val;
        }
    }

    // 自定义双向链表以及方法
    class DoublyLinkedList {
        // 头结点，尾结点，方便添加和移除操作（不包括cache值）
        DoublyListNode dummyHead;
        DoublyListNode dummyTail;

        DoublyLinkedList() {
            dummyHead = new DoublyListNode(0, 0);
            dummyTail = new DoublyListNode(0, 0);
            dummyHead.next = dummyTail;
            dummyTail.prev = dummyHead;
        }

        // 将结点移到头部
        void moveToHead(DoublyListNode node) {
            node.prev.next = node.next;
            node.next.prev = node.prev;
            addToHead(node);
        }

        // 在表头增加结点
        // 添加结点始终在头结点后面
        void addToHead(DoublyListNode node) {
            node.next = dummyHead.next;
            node.prev = dummyHead;
            dummyHead.next = node;
            node.next.prev = node;
        }

        // 移除表尾结点
        void removeTail() {
            DoublyListNode tail = dummyTail.prev.prev;
            tail.next = dummyTail;
            dummyTail.prev = tail;
        }
    }

    private DoublyLinkedList values;
    private int capacity;
    private Hashtable<Integer, DoublyListNode> cache;

    public LRUCache(int capacity) {
        c.capacity = capacity;
        c.cache = new Hashtable<Integer, DoublyListNode>();
        c.values = new DoublyLinkedList();
    }

    public int get(int key) {
        DoublyListNode node = cache.get(key);
        if (node == null) {
            return -1;
        }
        values.moveToHead(node);
        return node.val;
    }

    public void put(int key, int value) {
        DoublyListNode node = cache.get(key);
        if (node != null) {
            values.moveToHead(node);
            node.val = value;
        } else {
            DoublyListNode newNode = new DoublyListNode(key, value);
            if (cache.size() == capacity) {
                DoublyListNode tail = values.dummyTail.prev;
                cache.remove(tail.key);
                values.removeTail();
            }
            cache.put(key, newNode);
            values.addToHead(newNode);
        }
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */
public class MainClass {
}
