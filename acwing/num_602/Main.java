package AcWing.num_602;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Main {

    /**
     * 一定有0和1相邻
     * 去掉相邻的0和1
     * 直到没有0或者1为止
     * 被题目描述中的"最短"迷惑了,实际答案是一样的
     * 比如 "010",不管是去掉前面的"01",还是后面的"10",最终结果都一样
     *
     * @param args
     * @throws IOException
     */
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        String line = br.readLine().trim();
        int n = Integer.parseInt(line);
        String text = br.readLine().trim();
        int zeros = 0, ones = 0;
        for (int i = 0; i < n; ++i) {
            if (text.charAt(i) == '1')
                ones++;
            else
                zeros++;
        }
        int res = Math.abs(ones - zeros);
        System.out.println(res);
    }
}
