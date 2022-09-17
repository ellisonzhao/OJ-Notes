package AcWing.num_601;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Arrays;

public class Main {
    private static int[] stringToIntegerArray(String input) {
        if (input == null || input.length() == 0)
            return new int[0];
        String[] parts = input.trim().split(" ");
        int[] output = new int[parts.length];
        for (int index = 0; index < parts.length; index++) {
            String part = parts[index].trim();
            output[index] = Integer.parseInt(part);
        }
        return output;
    }

    private static int m;
    private static int n;
    private static int[] coins;

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int[] firstLine = stringToIntegerArray(br.readLine());
        m = firstLine[0];
        n = firstLine[1];
        coins = new int[n + 1];
        for (int i = 0; i < n; ++i) {
            String part = br.readLine().trim();
            coins[i] = Integer.parseInt(part);
        }
        //　细节处理,　
        Arrays.sort(coins, 0, n - 1);
        while (n > 0 && coins[n - 1] > m) {
            --n;
        }
        // 要凑出[1,m]区间内所有值
        //保证n之前都是小于等于m的硬币
        coins[n] = m + 1;
        if (coins[0] != 1) {
            System.out.println(-1);
        } else {
            int values = 0;
            int sum = 0;
            for (int i = 0; i < n; ++i) {
                // 用尽可能多的coins[i]凑到接近coins[i+1]
                // 如果要凑coins[i]只需要１个硬币,所以考虑凑到coins[i+1]-1
                // sum: 已有的硬币数量凑出的值
                // sum + coins[i]*k >= coins[i+1]-1
                // 于是得到k: coins[i+1]-1-sum 对coins[i]上取整
                int k = (coins[i + 1] - 1 - sum + coins[i] - 1) / coins[i];
                values += k;
                sum += k * coins[i];
            }
            System.out.println(values);
        }

    }
}
