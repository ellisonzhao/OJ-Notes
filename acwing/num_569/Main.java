package AcWing.num_569;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

/**
 * 计算C(n,m)：
 * 分子分解成质因数乘积
 * 分母分解成质因数乘积
 * 然后分子上每个质因数的个数减去分母上每个质因数的个数, 最后相乘
 * 例如　C(6,3)
 * 分子是　1*2*3*4*5*6 = 1*2*3*(2*2)*5*(2*3) = 1*(2*2*2*2)*(3*3)*5
 * 分母是　(1*2*3)*(1*2*3) = 1*(2*2)*(3*3)
 * 质因数只剩下2个2,1个5
 * C(6,3) = (2*2)*5
 */
public class Main {
    // 整数字符串转成整数数组
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

    private static int MOD = 1_000_000_000 + 7;
    private static int N = 2010;
    private static int COUNT = 0;// 统计1~n有多少个素数
    private static int[] primes = new int[N]; // 记录1~n 的素数
    private static int[] powers = new int[N];
    private static boolean[] state = new boolean[N]; // 标记是否为素数

    private static void get_primes(int n) {
        //　标记法统计素数
        for (int i = 2; i <= n; i++) {
            if (!state[i]) {
                // state[i]为false, 则说明这个i是素数
                primes[COUNT++] = i;
                for (int j = i * 2; j <= n; j += i) {
                    state[j] = true;
                }
            }
        }
    }

    private static int getCountOfPrime(int n, int p) {
        // 统计 n!包含多少个素因子p
        int cnt = 0;
        while (n > 0) {
            cnt += n / p;
            n /= p;
        }
        return cnt;
    }

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int[] firstLine = stringToIntegerArray(br.readLine());
        int n = firstLine[0];
        int s = firstLine[1];
        if (s > n) {
            System.out.println(0);
        } else {
            get_primes(n);
            for (int i = 0; i < COUNT; ++i) {
                // c(n,s)　中包含多少个质因子p
                int p = primes[i];
                powers[i] += getCountOfPrime(n, p) - getCountOfPrime(s, p) - getCountOfPrime(n - s, p);
            }
            long values = 1;
            for (int i = 0; i < COUNT; ++i) {
                int p = primes[i];
                while (powers[i]-- > 0)
                    values = values * p % MOD;
            }
            for (int i = 0; i < n - s; i++) {
                values = (values << 1) % MOD;
            }
            System.out.println(values);
        }
    }
}
