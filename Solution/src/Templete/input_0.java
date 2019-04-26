package Templete;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

/**
 * @author charlsonz
 * @date 2019-04-23 22:46
 */
public class input_0 {
    private static int[] stringToIntegerArray(String input) {
        String[] parts = input.trim().split(" ");
        int[] output = new int[parts.length];
        for (int index = 0; index < parts.length; index++) {
            String part = parts[index].trim();
            output[index] = Integer.parseInt(part);
        }
        return output;
    }

    // 固定N行输入
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int N = Integer.parseInt(br.readLine().trim());
        while (N-- > 0) {
            String line = br.readLine();
            if (line == null || line.length() == 0)
                break;
            int[] parts = stringToIntegerArray(line);
            for (int part : parts)
                System.out.println(part);
        }
    }
}
