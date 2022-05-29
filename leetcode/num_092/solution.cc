class Solution {
 public:
  ListNode *reverseBetween(ListNode *head, int left, int right) {
    ListNode *dummy = new ListNode(0, head);
    // 找到 left 位置节点
    ListNode *prevTail = dummy, *subTail = head;
    for (int i = 1; i < left; ++i) {
      prevTail = subTail;
      subTail = subTail->next;
    }

    // subTail 是 left 位置的节点，也是中间部分逆置后的表尾节点
    // prevTail 指向 left 之前部分的尾结点，也相当于 left--right
    // 中间部分的头结点
    for (int i = left + 1; i <= right; ++i) {
      // 当前要进行逆置操作的节点
      ListNode *curr = subTail->next;
      // 连接后面可能还要进行操作的部分；如果当前操作的是 right
      // 位置的节点，则直接连接到最后
      subTail->next = curr->next;
      curr->next = prevTail->next;
      prevTail->next = curr;
    }

    head = dummy->next;
    delete dummy;
    return head;
  }
};
