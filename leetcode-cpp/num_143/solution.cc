class Solution {
 public:
  void reorderList(ListNode *head) {
    if (!head || !(head->next) || !(head->next->next)) {
      return;
    }

    // 1. 链表等分成两个子链表
    ListNode *mid = findMiddle(head);
    ListNode *firstList = head;
    ListNode *secondList = mid->next;
    // 断开中间节点
    mid->next = nullptr;

    // 2. 后半部分逆置
    secondList = reverse(secondList);

    // 3. 合并链表
    // 注意，head 没有变化，仍是指向 firstList 的头结点
    mergeList(firstList, secondList);
  }

  // 找到链表的中点
  ListNode *findMiddle(ListNode *head) {
    if (!head || !(head->next)) {
      return head;
    }
    ListNode *slow = head, *fast = head;
    while (fast->next && fast->next->next) {
      fast = fast->next->next;
      slow = slow->next;
    }

    return slow;
  }

  ListNode *reverse(ListNode *head) {
    if (!head || !(head->next)) {
      return head;
    }
    ListNode *prev = nullptr;
    ListNode *curr = head;
    while (curr) {
      ListNode *next = curr->next;
      curr->next = prev;

      prev = curr;
      curr = next;
    }

    return prev;
  }

  void mergeList(ListNode *listA, ListNode *listB) {
    while (listA && listB) {
      ListNode *a = listA->next;
      ListNode *b = listB->next;

      listA->next = listB;
      listA = a;
      listB->next = listA;
      listB = b;
    }
  }
};