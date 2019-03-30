package nowcoder.toutiao_2;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

/**
 * @author charlsonz
 * @date 2019-03-29 19:21
 */
public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        Scanner sc = new Scanner(System.in);
        int T = sc.nextInt();
        for (int t = 0; t < T; t++) {
            String input = br.readLine();
            input.trim();
            // 数组：字符串数字
            String[] parts = input.split(" ");
            int N = parts.length; // 原始有多少张牌
            int[] sequence = new int[2 * N - 1]; // 手中牌序列
            for (int i = 0; i < N; i++) {
                sequence[i] = i;
            }
            int start = 0, end = N;
            while (start < sequence.length) {
                if (start % 2 != 0)
                    sequence[end++] = sequence[start];
                start++;
            }
            Map<Integer, Integer> dict = new HashMap<>();
            for (int i = 0; i < N; i++) {
                dict.put(sequence[2 * i], Integer.parseInt(parts[i]));
            }
            StringBuilder sb = new StringBuilder();
            for (int i = 0; i < N; i++) {
                sb.append(dict.get(i));
                sb.append(" ");
            }
            sb.deleteCharAt(sb.length() - 1);
            System.out.println(sb.toString());
        }

    }
}