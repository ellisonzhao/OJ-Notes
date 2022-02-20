package AcWing.num_570;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Collections;

/**
 * 利用最大滑动窗口问题求解
 * 参考　https://leetcode.com/problems/minimum-window-substring/
 */
public class Main {
    private static int[] stringToIntegerArray(String input) {
        input = input.trim();
        if (input.length() == 0) {
            return new int[0];
        }
        String[] parts = input.split(" ");
        int[] output = new int[parts.length];
        for (int index = 0; index < parts.length; index++) {
            String part = parts[index].trim();
            output[index] = Integer.parseInt(part);
        }
        return output;
    }

    private static int minWindow(int[] nums, int m) {
        int counter = m;// 用于判断窗口中是否包含所有编号
        int[] dict = new int[m + 1]; // 记录编号出现次数
        for (int i = 1; i <= m; ++i) {
            ++dict[i];
        }
        int begin = 0, end = 0;// 双指针,记录滑动窗口的起点和终点
        int width = Integer.MAX_VALUE;
        while (end < nums.length) {
            int endNum = nums[end++];
            if (dict[endNum]-- > 0)
                --counter;
            // 当前窗口已经包含所有编号
            while (counter == 0) {
                if (end - begin < width) {
                    width = end - begin;
                }
                int beginNum = nums[begin++];
                if (dict[beginNum]++ == 0) {
                    ++counter;
                }
            }
        }
        return width == Integer.MAX_VALUE ? -1 : width;
    }

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int[] firstLine = stringToIntegerArray(br.readLine());
        int n = firstLine[0];
        int m = firstLine[1];
        if (m > n) {
            System.out.println(-1);
        } else {
            int[] nums = stringToIntegerArray(br.readLine());
            int res = minWindow(nums, m);
            System.out.println(res);
        }
    }
}
