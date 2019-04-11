package AcWing.num_601;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Arrays;

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

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int[] firstLine = stringToIntegerArray(br.readLine());
        int m = firstLine[0];
        int n = firstLine[1];
        int[] nums = new int[n];
        for (int i = 0; i < n; ++i) {
            String part = br.readLine().trim();
            nums[i] = Integer.parseInt(part);
        }
        Arrays.sort(nums);

    }
}
