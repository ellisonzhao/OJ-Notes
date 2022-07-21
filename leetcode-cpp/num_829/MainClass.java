package LeetCode.num_829;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.List;

class Solution {
    public int consecutiveNumbersSum(int sum) {
        List<List<Integer>> res = new ArrayList<>();
        int k = ((sum & 1) == 0 ? 0 : 1) + sum / 2, curSum = 0;
        List<Integer> temp = new ArrayList<>();
        for (int i = 1; i <= k; i++) {
            temp.add(i);
            curSum += i;
            if (temp.size() < 2)
                continue;
            while (curSum > sum) {
                curSum -= temp.get(0);
                temp.remove(0);
            }
            if (curSum == sum) {
                res.add(new ArrayList<>(temp));
            }
        }
        return res.size();
    }
}

public class MainClass {
    public static void main(String[] args) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        String line;
        while ((line = in.readLine()) != null) {
            int N = Integer.parseInt(line);

            int ret = new Solution().consecutiveNumbersSum(N);

            String out = String.valueOf(ret);

            System.out.print(out);
        }
    }
}
