package temp.test01;

/**
 * @author charlsonz
 * @date 2019-04-10 21:58
 */

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Comparator;
import java.util.TreeSet;

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
        int n = Integer.parseInt(br.readLine().trim());
        int[] A = stringToIntegerArray(br.readLine());
        /**
         * AcWing似乎不支持Java里面的 javafx.util.Pair 包,因此用长度为2的数组来代替
         * 重写比较器防止值溢出
         * 如果直接用lambda表达式
         * TreeSet<int[]> treeSet = new TreeSet<>((o1, o2) -> o1[0] - o2[0]);
         * 那么返回o1[0]-o2[0]时可能会造成值溢出
         */
        TreeSet<int[]> treeSet = new TreeSet<>(new Comparator<int[]>() {
            @Override
            public int compare(int[] o1, int[] o2) {
                if (o1[0] < o2[0])
                    return -1;
                else if (o1[0] > o2[0])
                    return 1;
                else return 0;
            }
        });
        treeSet.add(new int[]{Integer.MIN_VALUE, 0});
        treeSet.add(new int[]{A[0], 1});
        treeSet.add(new int[]{Integer.MAX_VALUE, 0});

        for (int i = 1; i < n; ++i) {
            // 第i+1个数
            int[] pair = new int[]{A[i], i + 1};

            // 小于等于A[i]的最大值和A[i]的距离
            int[] leftPair = treeSet.floor(pair);
            long leftVal = A[i] - (long) leftPair[0];

            // 大于等于A[i]的最小值和A[i]的距离
            int[] rightPair = treeSet.ceiling(pair);
            long rightVal = (long) rightPair[0] - A[i];

            // A[i]距离谁更近
            if (leftVal <= rightVal) {
                // 距离小于等于A[i]的最大值更近
                System.out.println(leftVal + " " + leftPair[1]);
            } else {
                System.out.println(rightVal + " " + rightPair[1]);
            }
            treeSet.add(pair);
        }
    }
}
