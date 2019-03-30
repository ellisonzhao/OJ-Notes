package nowcoder.monthlyTest3_2;

/**
 * @author charlsonz
 * @date 2019-03-14 20:04
 */
// 本题为考试多行输入输出规范示例，无需提交，不计分。

import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
//        int ans = 0, x;
        long[][] data = new long[n][3];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < 3; j++) {
                data[i][j] = sc.nextInt();
            }
        }

//        for(int[] a : data)
//            System.out.println(a.toString());
    }
}
