package LeetCode.num_17;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

class Solution {
    public List<String> letterCombinations(String digits) {
        List<String> ans = new ArrayList<>();
        if (digits == null || digits.length() == 0) {
            return ans;
        }
        String[] dict = new String[]{"0", "1", "abc", "def", "ghi",
                "jkl", "mno", "pqrs", "tuv", "wxyz"};
        Queue<String> q = new LinkedList<>();
        q.add("");
        for (int i = 0; i < digits.length(); i++) {
            int digit = digits.charAt(i) - '0';
            int size = q.size();
            for (int j = 0; j < size; j++) {
                String temp = q.poll();
                for (int k = 0; k < dict[digit].length(); k++) {
                    q.add(temp + dict[digit].charAt(k));
                }
            }
        }
        while (!q.isEmpty()) {
            ans.add(q.peek());
            q.poll();
        }
        return ans;
    }
}

public class MainClass {
}
