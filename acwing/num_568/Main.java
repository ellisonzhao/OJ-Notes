package AcWing.num_568;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

/**
 * @author charlsonz
 * @date 2019-04-02 19:19
 */

public class Main {

    private static int getIntervalSum(int l, int r) {
        int k = (r - l + 1) / 2;
        int result = (l % 2 == 0 ? -k : k);
        if ((r - l + 1) % 2 == 1) {
            if (r % 2 == 0) {
                result += r;
            } else {
                result -= r;
            }
        }
        return result;
    }

    public static void main(String[] args) throws IOException {
        long startTime = System.currentTimeMillis(); //获取开始时间

        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int q = Integer.valueOf(br.readLine());
        while (q-- > 0) {
            String[] parts = br.readLine().split(" ");
            int l = Integer.valueOf(parts[0]);
            int r = Integer.valueOf(parts[1]);
            System.out.println(getIntervalSum(l, r));
        }

        long endTime = System.currentTimeMillis(); //获取结束时间
        System.out.println("程序运行时间：" + (endTime - startTime) + "ms"); //输出程序运行时间
    }
}
