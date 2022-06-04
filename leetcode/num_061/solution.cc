class Solution {
 public:
  ListNode *rotateRight(ListNode *head, int k) {
    if (!head || !k) return head;
    // 1. 计算链表长度，并找到尾结点
    int n = 0;       // 链表的长度
    ListNode *tail;  // 尾节点
    for (ListNode *p = head; p; p = p->next) {
      tail = p;
      n++;
    }
    // 2. 找到头结点旋转后的位置
    k %= n;
    ListNode *p = head;
    for (int i = 0; i < n - k - 1; i++) p = p->next;  // 找到链表的第n-k个节点
    // 3. 旋转拼接
    tail->next = head;
    head = p->next;
    p->next = nullptr;  // 链表成环，注意断开
    return head;        // 返回新的头节点
  }
};