package LeetCode.Weekly_Contest_107;

class Solution {
    public boolean isLongPressedName(String name, String typed) {
        int n = name.length(), t = typed.length();
        int[] dict = new int[n];
        char[] cname = name.toCharArray(), ctype = typed.toCharArray();
        dict[0] = 1;
        int count = 0;
        for (int i = 1; i < n; ++i) {
            if (cname[i] == cname[i - 1]) {
                dict[count]++;
            } else {
                count++;
                dict[count] = 1;
            }
        }

        count = 0;
        dict[0]--;
        for (int i = 1; i < t; ++i) {
            if (ctype[i] != ctype[i - 1])
                count++;
            dict[count]--;
        }

        for (int d : dict) {
            if (d > 0)
                return false;
        }

        return true;
    }
}

public class MainClass {
    public static void main(String[] args) {
        String name = "alex", typed = "aaleex";
        System.out.println(new Solution().isLongPressedName(name, typed));
    }
}
