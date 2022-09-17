package LeetCode.num_967;


import java.util.*;

class Solution {
    public int[] numsSameConsecDiff(int N, int K) {
        Queue<String> queue = new LinkedList<>();
        Set<String> values = new HashSet<>();
        for (int i = 0; i <= 9; ++i) {
            queue.add(String.valueOf(i));
        }
        while (!queue.isEmpty()) {
            String temp = queue.poll();
            if (temp.length() == N) {
                values.add(temp);
                continue;
            } else {
                char first = temp.charAt(0);
                if (first == '0') {
                    continue;
                }
                char c = temp.charAt(temp.length() - 1);
                int lastDegit = c - '0';
                if (lastDegit - K >= 0 && lastDegit - K <= 9) {
                    queue.add(temp + (lastDegit - K));
                }
                if (lastDegit + K >= 0 && lastDegit + K <= 9) {
                    queue.add(temp + (lastDegit + K));
                }

            }
        }

        List<String> values = new ArrayList<>(values);
        int[] ans = new int[values.size()];
        for (int i = 0; i < values.size(); i++) {
            ans[i] = Integer.valueOf(values.get(i));
        }
        return ans;
    }

}

public class MainClass {
    public static void main(String[] args) {
        int ans[] = new Solution().numsSameConsecDiff(2, 0);
        for (int arr : ans) {
            System.out.println(arr);
        }
    }
}
