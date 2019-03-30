package nowcoder.toutiao_1;

import java.util.Arrays;
import java.util.Scanner;

/**
 * @author charlsonz
 * @date 2019-03-16 10:45
 */
public class Main {
    private void getFrontReward(int[] arr, int[] res, int n, int index) {
        int pointer = index;
        int next = (pointer + 1) % n;
        while (next != index) {
            if (arr[next] > arr[pointer]) {
                res[next] = Math.max(res[next], res[pointer] + 1);
            }
            next = (next + 1) % n;
        }
    }

    private void getBackReward(int[] arr, int[] res, int n, int index) {
        int pointer = index;
        int next = (pointer + n - 1) % n;
        while (next != index) {
            if (arr[next] > arr[pointer]) {
                res[next] = Math.max(res[next], res[pointer] + 1);
            }
            next = (next - 1 + n) % n;
        }
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int T = sc.nextInt();
        for (int i = 0; i < T; i++) {
            int n = sc.nextInt();
            int[] arr = new int[n];
            for (int j = 0; j < n; j++) {
                arr[j] = sc.nextInt();
            }
//            for(int a:arr){
//                System.out.println(a);
//            }
            int mini = Integer.MAX_VALUE;
            int index = -1;
            for (int k = 0; k < n; k++) {
                if (arr[k] < mini) {
                    index = k;
                    mini = arr[k];
                }
            }

            int[] res = new int[n];
            Arrays.fill(res, 1);
            Main obj = new Main();
            obj.getFrontReward(arr, res, n, index);
            obj.getBackReward(arr, res, n, index);

            int ans = 0;
            for (int a : res) {
                ans += a;
            }
            System.out.println(ans);
        }
    }

}
