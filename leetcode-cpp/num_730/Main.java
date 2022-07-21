package LeetCode.num_730;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

/**
 * @author charlsonz
 * @date 2019-04-24 09:40
 */
public class Main {
    private static int[] stringToIntegerArray(String input) {
        String[] parts = input.trim().split(" ");
        int[] output = new int[parts.length];
        for (int index = 0; index < parts.length; index++) {
            String part = parts[index].trim();
            output[index] = Integer.parseInt(part);
        }
        return output;
    }

    private static boolean check(long energy) {
        for (int i = 1; i <= N; ++i) {
            //
            energy = energy - (H[i - 1] - energy);
            if (energy < 0)
                return false;
            // 达到阈值肯定可以到达最后建筑
            if (energy >= 1e10)
                return true;
        }
        return true;
    }

    private static int N;
    private static int[] H;

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        N = Integer.parseInt(br.readLine().trim());
        H = stringToIntegerArray(br.readLine());

        long l = 0, r = (long) 1e10; // 结果返回整数
//        System.out.println(r);
        while (l < r) {
            long mid = (l + r) >>> 1;
            if (check(mid)) {
                r = mid;
            } else {
                l = mid + 1;
            }
        }
        System.out.println(l);
    }
}
