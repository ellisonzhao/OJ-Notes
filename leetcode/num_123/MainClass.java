package LeetCode.num_123;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

class Solution {
    public int maxProfit(int[] prices) {
        int hold1 = Integer.MIN_VALUE, hold2 = Integer.MIN_VALUE;
        int release1 = 0, release2 = 0;
        for (int price : prices) {                              // Assume we only have 0 money at first
            release2 = Math.max(release2, hold2 + price);     // The maximum if we've just sold 2nd stock so far.
            hold2 = Math.max(hold2, release1 - price);  // The maximum if we've just buy  2nd stock so far.
            release1 = Math.max(release1, hold1 + price);     // The maximum if we've just sold 1nd stock so far.
            hold1 = Math.max(hold1, -price);          // The maximum if we've just buy  1st stock so far.
        }
        return release2; ///Since release1 is initiated as 0, so release2 will always higher than release1.
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
            int[] prices = stringToIntegerArray(line);

            int ret = new Solution().maxProfit(prices);

            String out = String.valueOf(ret);

            System.out.print(out);
        }
    }
}