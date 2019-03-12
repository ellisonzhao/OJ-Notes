package LeetCode.num_784;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

class Solution {

    public List<String> letterCasePermutation(String S) {
        if (S == null) {
            return new LinkedList<>();
        }
        Queue<String> queue = new LinkedList<>();
        queue.offer(S);

        for (int i = 0; i < S.length(); ++i) {
            if (Character.isDigit(S.charAt(i)))
                continue;
            int size = queue.size();
            for (int j = 0; j < size; j++) {
                String cur = queue.poll();
                char[] chs = cur.toCharArray();

                chs[i] = Character.toUpperCase(chs[i]);
                queue.offer(String.valueOf(chs));

                chs[i] = Character.toLowerCase(chs[i]);
                queue.offer(String.valueOf(chs));
            }
        }

        return new LinkedList<>(queue);
    }

    public boolean detectCapitalUse(String word) {
        for (int i = 1; i < word.length(); ++i) {
            if ((Character.isLowerCase(word.charAt(1)) !=
                    Character.isLowerCase(word.charAt(i))) ||
                    (Character.isLowerCase(word.charAt(0)) &&
                            Character.isUpperCase(word.charAt(i))))
                return false;
        }
        return true;
    }

    public int countBinarySubstrings(String S) {
        int ans = 0, N = S.length(), index = 0;
        int[] group = new int[N];
        group[0] = 1;
        for (int i = 1; i < N; ++i) {
            if (S.charAt(i) != S.charAt(i - 1)) {
                group[++index] = 1;
            } else {
                ++group[index];
            }
        }

        //index此时为最后一个分组的下标
        for (int i = 1; i <= index; ++i) {
            ans += Math.min(group[i - 1], group[i]);
        }
        return ans;
    }

    public static void main(String[] args) {
        String S = "leeTcode";
        System.out.println(new Solution().detectCapitalUse(S));
    }
}
