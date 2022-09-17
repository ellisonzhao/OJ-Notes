package LeetCode.num_206;

/* -----------------------------------
 *  WARNING:
 * -----------------------------------
 *  Your code may fail to compile
 *  because it contains public class
 *  declarations.
 *  To fix c, please remove the
 *  "public" keyword from your class
 *  declarations.
 */

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

/**
 * Definition for singly-linked values.
 * public class ListNode {
 * int val;
 * ListNode next;
 * ListNode(int x) { val = x; }
 * }
 */
class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }

}

class Solution {
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null)
            return head;
        ListNode L = new ListNode(-1);
        ListNode p = head, queue = null;
        while (p != null) {
            queue = p.next;
            p.next = L.next;
            L.next = p;
            p = queue;
        }
        return L.next;
    }

    public ListNode reverseBetween(ListNode head, int m, int n) {
        if (head == null || head.next == null)
            return head;
        ListNode L = new ListNode(-1), cur = head, rear = null;
        L.next = cur;
        head = L;
        int index = 1; // record the index of cur Node
        while (index < m) {
            cur = cur.next;
            head = head.next;
            ++index;
        }
        rear = cur;
        cur = cur.next;
        ++index;
        while (index <= n) {
            ListNode temp = cur.next;
            cur.next = head.next;
            head.next = cur;
            cur = temp;
            ++index;
        }
        rear.next = cur;

        return L.next;
    }

    public double lengthOfList(ListNode head) {
        int length = 0;
        while (head != null) {
            ++length;
            head = head.next;
        }
        return length;
    }

    public boolean isPalindrome(ListNode head) {
        if (head == null || head.next == null)
            return true;

        double length = lengthOfList(head);
        int mid = (int) Math.ceil(length / 2);
        int index = 1;
        ListNode middle = head;
        while (index < mid) {
            middle = middle.next;
            ++index;
        }
        ListNode front = reverseList(middle.next);
        while (front != null) {
            if (front.val != head.val)
                return false;
            front = front.next;
            head = head.next;
        }
        return true;
    }
}

public class MainClass {
    public static int[] stringToIntegerArray(String input) {
        input = input.trim();
        input = input.substring(1, input.length() - 1);
        if (input.length() == 0) {
            return new int[0];
        }

        String[] parts = input.split(",");
        int[] output = new int[parts.length];
        for (int index = 0; index < parts.length; index++) {
            String part = parts[index].trim();
            output[index] = Integer.parseInt(part);
        }
        return output;
    }

    public static ListNode stringToListNode(String input) {
        // Generate array from the input
        int[] nodeValues = stringToIntegerArray(input);

        // Now convert that values into linked values
        ListNode dummyRoot = new ListNode(0);
        ListNode ptr = dummyRoot;
        for (int item : nodeValues) {
            ptr.next = new ListNode(item);
            ptr = ptr.next;
        }
        return dummyRoot.next;
    }

    public static String listNodeToString(ListNode node) {
        if (node == null) {
            return "[]";
        }

        String result = "";
        while (node != null) {
            result += Integer.toString(node.val) + ", ";
            node = node.next;
        }
        return "[" + result.substring(0, result.length() - 2) + "]";
    }

    public static void main(String[] args) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        String line;
        while ((line = in.readLine()) != null) {
            ListNode head = stringToListNode(line);

//            ListNode ret = new Solution().reverseBetween(head, 2, 4);

            boolean ret = new Solution().isPalindrome(head);

//            String out = listNodeToString(ret);

//            System.out.print(out);

            System.out.println(ret);
        }
    }
}
