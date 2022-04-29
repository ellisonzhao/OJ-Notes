/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
 public:
  ListNode *removeNthFromEnd(ListNode *head, int n) {
    ListNode *dummy = new ListNode(0, head);
    ListNode *fast = dummy;
    ListNode *slow = dummy;
    while (n) {
      fast = fast->next;
      n--;
    }
    while (fast->next) {
      fast = fast->next;
      slow = slow->next;
    }
    auto temp = slow->next;
    slow->next = temp->next;
    delete temp;
    head = dummy->next;
    delete dummy;
    return head;
  }
};