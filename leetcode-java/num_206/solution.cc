class Solution {
 public:
  ListNode *reverseList(ListNode *head) {
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
};