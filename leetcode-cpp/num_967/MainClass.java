package LeetCode.num_967;


import java.util.*;

class Solution {
    public int[] numsSameConsecDiff(int N, int K) {
        Queue<String> q = new LinkedList<>();
        Set<String> res = new HashSet<>();
        for (int i = 0; i <= 9; ++i) {
            q.add(String.valueOf(i));
        }
        while (!q.isEmpty()) {
            String temp = q.poll();
            if (temp.length() == N) {
                res.add(temp);
                continue;
            } else {
                char first = temp.charAt(0);
                if (first == '0') {
                    continue;
                }
                char c = temp.charAt(temp.length() - 1);
                int lastDegit = c - '0';
                if (lastDegit - K >= 0 && lastDegit - K <= 9) {
                    q.add(temp + (lastDegit - K));
                }
                if (lastDegit + K >= 0 && lastDegit + K <= 9) {
                    q.add(temp + (lastDegit + K));
                }

            }
        }

        List<String> list = new ArrayList<>(res);
        int[] ans = new int[list.size()];
        for (int i = 0; i < list.size(); i++) {
            ans[i] = Integer.valueOf(list.get(i));
        }
        return ans;
    }

}

public class MainClass {
    public static void main(String[] args) {
        int ans[] = new Solution().numsSameConsecDiff(2, 0);
        for (int a : ans) {
            System.out.println(a);
        }
    }
}
