package AcWing.num_680;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

/**
 * @author charlsonz
 * @date 2019-04-15 21:07
 */
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

    private static boolean check(double mid) {
        int sum = 0;
        for (int l : L) {
            sum += l / mid;
            if (sum >= M) {
                return true;
            }
        }
        return false;
    }

    private static int M;
    private static int[] L;

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int[] firstLine = stringToIntegerArray(br.readLine());
        M = firstLine[1];
        L = stringToIntegerArray(br.readLine().trim());
        double l = 0, r = 1e9;
        while (r - l > 1e-6) {
            double mid = (l + r) / 2;
            if (check(mid))
                l = mid;
            else
                r = mid;
        }
        System.out.println(String.format("%.2f", l));
    }
}
