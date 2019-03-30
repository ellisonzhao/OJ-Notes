package LeetCode.num_301;

import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

/**
 * @author charlsonz
 * @date 2019-03-22 21:57
 */

class Solution {
    public List<String> removeInvalidParentheses(String s) {
        List<String> res = new ArrayList<>();
        if (s == null)
            return res;
        remove(s, res, 0, 0, new char[]{'(', ')'});
        return res;
    }

    private void remove(String s, List<String> res, int last_invalid, int last_remove, char[] brackets) {
        for (int counter = 0, i = last_invalid; i < s.length(); i++) {
            if (s.charAt(i) == brackets[0])
                counter++;
            if (s.charAt(i) == brackets[1])
                counter--;
            if (counter >= 0)
                continue;
            // 右括号较多，每次移除移除第一个
            // 如果是左括号较多，逆置字符串之后，变成')' 与 '(' 配对
            // 这样就是删除左括号较多的情况
            for (int j = last_remove; j <= i; j++) {
                // 第一个不匹配的地方
                // 因为可能还有其他字符，所以条件就是不等于 brackets[1]
                if (s.charAt(j) == brackets[1] && (j == last_remove || s.charAt(j - 1) != brackets[1]))
                    remove(s.substring(0, j) + s.substring(j + 1), res, i, j, brackets);
            }
            return;
        }
        String reversed = new StringBuilder(s).reverse().toString();
        if (brackets[0] == '(') {
            // 完成从左到右对右括号的扫描
            // 继续对左括号扫描
            remove(reversed, res, 0, 0, new char[]{')', '('});
        } else {
            // 全部扫描完成
            res.add(reversed);
        }
    }

    public int largestRectangleArea(int[] heights) {
        int len = heights.length;
        int maxArea = 0;
        Stack<Integer> stack = new Stack<>();
        for (int i = 0; i <= len; ) {
            int h = (i == len ? 0 : heights[i]);
            if (stack.isEmpty() || h >= heights[stack.peek()]) {
                stack.push(i);
                i++;
            } else {
                int curHeight = heights[stack.pop()];
                int rightBoundary = i - 1;
                int leftBoundary = stack.isEmpty() ? 0 : stack.peek() + 1;
                int width = rightBoundary - leftBoundary + 1;
                maxArea = Math.max(maxArea, (curHeight * width));
            }
        }
        return maxArea;
    }
}

public class MainClass {
    public static void main(String[] args) {
        String s = "()())()";
        System.out.println(new Solution().removeInvalidParentheses(s));
    }
}
