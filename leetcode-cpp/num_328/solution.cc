class Solution {
public:
  ListNode *oddEvenList(ListNode *head) {
    if (!head)
      return head;

    ListNode *odd = head;
    ListNode *evenHead = head->next, *even = evenHead;
    while (even && even->next) {
      odd->next = even->next;
      odd = odd->next;
      even->next = even->next->next;
      even = even->next;
    }
    odd->next = evenHead;

    return head;
  }
};