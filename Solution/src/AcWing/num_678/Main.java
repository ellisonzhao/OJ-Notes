package AcWing.num_678;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

/**
 * @author charlsonz
 * @date 2019-04-15 15:34
 */
public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int T = Integer.parseInt(br.readLine().trim());
        while (T-- > 0) {
            char[] src = br.readLine().trim().toCharArray();
            char[] target = new char[src.length];
            int j = 0;
            // j 就是转换后字符串的长度
            for (int i = 0; i < src.length; ++i) {
                target[j++] = src[i];
                if (j >= 3) {
                    // 遇到连续两个相同字符
                    boolean flag = target[j - 1] == target[j - 2];

                    // 三个相同字符: helllo 式
                    if ((flag && target[j - 2] == target[j - 3]) ||
                            // 或者AABB式
                            (j >= 4 && flag && target[j - 3] == target[j - 4])
                    ) {
                        --j;
                    }
                }
            }

            System.out.println(String.valueOf(target, 0, j));
        }
    }
}
