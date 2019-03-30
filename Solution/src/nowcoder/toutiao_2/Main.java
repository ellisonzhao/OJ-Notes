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
class Solution {
    public int[] getOriginCardSequence(int[] inputSequence) {
        int N = inputSequence.length; // 原始有多少张牌
        if (N == 0)
            return new int[0];
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
            dict.put(sequence[2 * i], inputSequence[i]);
        }
        int[] result = new int[N];
        for (int i = 0; i < N; i++) {
            result[i] = dict.get(i);
        }
        return result;
    }
}

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

    private static String integerArrayToString(int[] nums) {
        if (nums.length == 0) {
            return "";
        }
        StringBuilder sb = new StringBuilder();
        for (int num : nums) {
            sb.append(num);
            sb.append(" ");
        }
        String output = sb.toString();
        output = output.trim();
        return output;
    }

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        Scanner sc = new Scanner(System.in);
        int T = sc.nextInt();
        for (int t = 0; t < T; t++) {
            String line = br.readLine();
            int[] inputSequence = stringToIntegerArray(line);
            int[] result = new Solution().getOriginCardSequence(inputSequence);
            System.out.println(integerArrayToString(result));
        }

    }
}