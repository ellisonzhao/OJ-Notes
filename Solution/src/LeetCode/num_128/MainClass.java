package LeetCode.num_128;

import java.util.HashSet;
import java.util.Set;

class Solution {
    public int longestConsecutive(int[] nums) {
        if (nums == null || nums.length == 0)
            return 0;
        Set<Integer> set = new HashSet<>();
        for (int num : nums) {
            set.add(num);
        }
        int ans = 0;
        for (int num : set) {
            if (!set.contains(num - 1)) {
                int start = num;
                int cur = 1;
                while (set.contains(start + 1)) {
                    start++;
                    cur++;
                }
                ans = Math.max(ans, cur);
            }
        }
        return ans;
    }
}

public class MainClass {
}
