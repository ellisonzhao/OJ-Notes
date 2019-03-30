package nowcoder.monthlyTest3_1;

/**
 * @author charlsonz
 * @date 2019-03-14 19:44
 */

import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class Main {
    private static Map<Integer, String> map = new HashMap<>();
    private static int res = Integer.MAX_VALUE;

    private void traverse(int id, int k, int n, int sum, String prefix) {
        if (k < 0) return;
        if (k == 0 && n == 0) {
            map.put(sum, prefix);
            res = Math.min(sum, res);
        }
        traverse(id * 2, k - 1, n - id, sum + id, prefix + id + " + ");
        traverse(id * 2 + 1, k - 1, n - id, sum + id, prefix + id + " + ");
        traverse(id * 2, k - 1, n + id, sum + id, prefix + id + " - ");
        traverse(id * 2 + 1, k - 1, n + id, sum + id, prefix + id + " - ");
    }

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        while (in.hasNextInt()) {//注意while处理多个case
            int n = in.nextInt();
            int k = in.nextInt();
            String str = "";
            Main obj = new Main();
            obj.traverse(1, k, n, 0, str);
            String ans = map.get(res);
            ans.trim();
            String[] result = ans.split(" ");
            int count = 0;
            for (int i = 0; i < result.length; i++) {
                if (count % 2 == 0)
                    System.out.print(result[i]);
                if (count % 2 == 1)
                    System.out.println(" " + result[i]);
                count++;
            }
            break;
        }

    }

}
