package AcWing.num_677;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int n = Integer.parseInt(br.readLine().trim());
        n = 1024 - n;
        int res = 0;

        int[] coins = {64, 16, 4, 1};
        for (int i = 0; i < 4; ++i) {
            res += n / coins[i];
            n = n % coins[i];
        }

        System.out.println(res);
    }
}
