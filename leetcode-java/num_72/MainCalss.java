package LeetCode.num_72;

import java.util.Arrays;

/**
 * @author charlsonz
 * @date 2019-03-29 13:35
 */
class Solution {
    private int[][] dp;

    public int minDistance(String word1, String word2) {
        if (word1 == null || word2 == null)
            return 0;
        int l1 = word1.length(), l2 = word2.length();

        dp = new int[l1 + 1][l2 + 1];
        for (int[] arr : dp)
            Arrays.fill(arr, -1);
        return minDistance(word1, word2, l1, l2);
    }

    private int minDistance(String word1, String word2, int l1, int l2) {
        if (l1 == 0) return l2;
        if (l2 == 0) return l1;
        if (dp[l1][l2] >= 0) return dp[l1][l2];
        int values;
        if (word1.charAt(l1 - 1) == word2.charAt(l2 - 1)) {
            values = minDistance(word1, word2, l1 - 1, l2 - 1);
        } else {
            // 插入一个字符:l1,l2-1
            // 删除最后一个字符:l1-1, l2
            // 替换最后一个字符: l1-1,l2-1
            values = 1 + Math.min(minDistance(word1, word2, l1 - 1, l2 - 1),
                    Math.min(minDistance(word1, word2, l1 - 1, l2), minDistance(word1, word2,
                            l1, l2 - 1)));
        }
        dp[l1][l2] = values;
        return values;
    }
}

public class MainCalss {
}
