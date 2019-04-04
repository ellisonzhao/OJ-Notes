package LeetCode.num_76;

class Solution {
    public String minWindow(String s, String t) {
        int[] dict = new int[128];
        for (char c : t.toCharArray())
            dict[c]++;
        int counter = t.length(); // 用于判断s中的子串是否包含t中所有字符
        int begin = 0, end = 0; // 双指针,记录滑动窗口的起点和终点
        int head = 0; //　记录窗口长度最小时的宽度
        int width = Integer.MAX_VALUE; //　窗口宽度,即子串长度
        // 这样根据　head 和　width　即可截取子串
        while (end < s.length()) {
            char endChar = s.charAt(end++);
            if (dict[endChar]-- > 0)
                --counter;
            while (counter == 0) {
                // 当前窗口已经包含所有字符
                if (end - begin < width) {
                    //　长度更小的窗口
                    head = begin;
                    width = end - begin;
                }

                //
                char beginChar = s.charAt(begin++);
                if (dict[beginChar]++ == 0) {
                    ++counter;
                }
            }
        }
        return width == Integer.MAX_VALUE ? "" : s.substring(head, head + width);
    }
}

public class MainClass {
    public static void main(String[] args) {
        String s = "ADOBECODEBANC", t = "ABC";
        System.out.println(new Solution().minWindow(s, t));
    }
}
