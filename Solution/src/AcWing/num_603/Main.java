package AcWing.num_603;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Arrays;

public class Main {
    private static long[] stringToLongArray(String input) {
        String[] parts = input.trim().split(" ");
        long[] output = new long[parts.length];
        for (int index = 0; index < parts.length; index++) {
            String part = parts[index].trim();
            output[index] = Long.parseLong(part);
        }
        return output;
    }

    private static int[] stringToIntegerArray(String input) {
        String[] parts = input.trim().split(" ");
        int[] output = new int[parts.length];
        for (int index = 0; index < parts.length; index++) {
            String part = parts[index].trim();
            output[index] = Integer.parseInt(part);
        }
        return output;
    }

    private static int N;
    private static long[] d;
    private static int[] p;

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        N = Integer.parseInt(br.readLine().trim());
        d = stringToLongArray(br.readLine());
        p = stringToIntegerArray(br.readLine());

        long[][] dp = new long[N + 1][N * 2 + 1];
        for (long[] arr : dp)
            Arrays.fill(arr, -1);
        dp[0][0] = 0;

        for (int i = 1; i <= N; ++i) {
            for (int j = 1; j <= N * 2; ++j) {
                // dp[i][j]　表示走到第i只怪兽,一共花费金币j能买到的最大武力值
                // 注意对应的武力值和价格是d[i-1],p[i-1]
                if (dp[i - 1][j] >= d[i - 1])
                    //　在当前怪兽处不需要花金币"贿赂"
                    dp[i][j] = dp[i - 1][j];
                if (j >= p[i - 1] && dp[i - 1][j - p[i - 1]] != -1) {
                    // 在当前怪兽处需要花费金币"贿赂"
                    //　并且在前面已经买过怪兽
                    dp[i][j] = Math.max(dp[i][j], dp[i - 1][j - p[i - 1]] + d[i - 1]);
                }
            }
        }

        int res = 0;
        for (int i = 1; i <= 2 * N; ++i) {
            if (dp[N][i] != -1) {
                res = i;
                break;
            }
        }
        System.out.println(res);
    }
}
