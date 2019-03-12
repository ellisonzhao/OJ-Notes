package LeetCode.num_53;


import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

class Solution {
    public int maxSubArray(int[] nums) {
        int n = nums.length;
        return maxSum(nums, 0, n);
    }

    //返回数组在左闭右开区间[x,y)中的最大连续和
    public int maxSum(int[] A, int x, int y) {
        int maxs;

        //只有一个元素直接返回
        if (y - x == 1) return A[x];

        //用(x+y)/2有可能出现值太大而越界的情形;
        //分治第一步：划分成[x,mid),[mid,y)
        int mid = x + (y - x) / 2;

        //分治第二步：递归求解
        maxs = Math.max(maxSum(A, x, mid), maxSum(A, mid, y));

        int v, L, R;
        //分治第三步：合并（1）————从分界点开始往左的最大连续和L
        v = 0;
        L = A[mid - 1];
        for (int i = mid - 1; i >= x; --i) {
            v += A[i];
            L = Math.max(L, v);
        }
        v = 0;
        R = A[mid];
        //分治第三步：合并（2）————从分界点开始往左的最大连续和R
        for (int i = mid; i < y; ++i) {
            v += A[i];
            R = Math.max(R, v);
        }
        //把子问题的解与L和R比较
        return Math.max(maxs, L + R);
    }
}

public class MainClass {
    public static int[] stringToIntegerArray(String input) {
        input = input.trim();
        input = input.substring(1, input.length() - 1);
        if (input.length() == 0) {
            return new int[0];
        }

        String[] parts = input.split(",");
        int[] output = new int[parts.length];
        for (int index = 0; index < parts.length; index++) {
            String part = parts[index].trim();
            output[index] = Integer.parseInt(part);
        }
        return output;
    }

    public static void main(String[] args) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        String line;
        while ((line = in.readLine()) != null) {
            int[] nums = stringToIntegerArray(line);

            int ret = new Solution().maxSubArray(nums);

            String out = String.valueOf(ret);

            System.out.print(out);
        }
    }
}