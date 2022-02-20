package LeetCode.num_394;

import java.util.Stack;

class Solution {
    public String decodeString(String s) {
        if (s == null || s.length() == 0)
            return s;
        Stack<String> resultStack = new Stack<>();
        Stack<Integer> timesStack = new Stack<>();
        String res = "";
        int index = 0;
        while (index < s.length()) {
            if (Character.isDigit(s.charAt(index))) {
                // 如果是数字，注意可能出现多位数
                int times = 0;
                while (Character.isDigit(s.charAt(index))) {
                    times = 10 * times + (s.charAt(index) - '0');
                    index++;
                }
                timesStack.push(times);
            } else if (s.charAt(index) == '[') {
                //前面的中间结果String入栈保存
                // 并且会有新的中间结果String
                resultStack.push(res);
                res = "";
                index++;
            } else if (s.charAt(index) == ']') {
                // 计算拼接字符串
                StringBuilder sb = new StringBuilder(resultStack.pop());
                int repeatTimes = timesStack.pop();
                for (int i = 0; i < repeatTimes; i++) {
                    sb.append(res);
                }
                res = sb.toString();
                index++;
            } else {
                // 中间结果String
                res += s.charAt(index);
                index++;
            }
        }
        return res;
    }
}

public class MainClass {
    public static void main(String[] args) {
        String s = "10[leetcode]";
        System.out.println(new Solution().decodeString(s));
    }
}
