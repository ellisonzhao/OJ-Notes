package AcWing.num_136;

import javafx.util.Pair;
import org.omg.PortableInterceptor.INACTIVE;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Comparator;
import java.util.Map;
import java.util.TreeMap;
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

//    public static void main(String[] args) throws IOException {
//        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
//        int n = Integer.parseInt(br.readLine().trim());
//        int[] A = stringToIntegerArray(br.readLine());
//        TreeSet<Pair<Integer, Integer>> treeSet = new TreeSet<>(Comparator.comparingInt(Pair::getKey));
//        treeSet.add(new Pair<>(Integer.MIN_VALUE, 0));
//        treeSet.add(new Pair<>(Integer.MAX_VALUE, 0));
//        for (int i = 0; i < n; ++i) {
//            Pair pair = new Pair<>(A[i], i + 1);
//            if (i > 0) {
//                // 大于等于A[i]的最小值
//                Pair lastPair = treeSet.ceiling(pair);
//                long rightVal = (int) lastPair.getKey() - A[i];
//
//                // 小于等于A[i]的最大值
//                Pair firstPair = treeSet.floor(pair);
//                long leftVal = A[i] - (int) firstPair.getKey();
//
//                // A[i]　距离谁更近
//                if (leftVal <= rightVal) {
//                    // 距离小于等于A[i]的最大值更近
//                    System.out.println(leftVal + " " + firstPair.getValue());
//                } else {
//                    System.out.println(rightVal + " " + lastPair.getValue());
//                }
//            }
//            treeSet.add(pair);
//        }
//
//    }

    // TreeMap
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int n = Integer.parseInt(br.readLine().trim());
        int[] A = stringToIntegerArray(br.readLine());
        TreeMap<Integer, Integer> treeMap = new TreeMap<>();
        treeMap.put(Integer.MIN_VALUE, 0);
        treeMap.put(A[0], 1);
        treeMap.put(Integer.MAX_VALUE, 0);

        for (int i = 1; i < n; ++i) {
            // 大于等于A[i]的最小值和A[i]的距离
            int rightKey = treeMap.ceilingKey(A[i]);
            long right_d = (long)rightKey - A[i];

            // 小于等于A[i]的最大值和A[i]的距离
            int leftKey = treeMap.floorKey(A[i]);
            long left_d = A[i] - (long)leftKey;

            // A[i]距离谁更近
            if (left_d <= right_d) {
                // 距离小于等于A[i]的最大值更近
                System.out.println(left_d + " " + treeMap.get(leftKey));
            } else {
                System.out.println(right_d + " " + treeMap.get(rightKey));
            }
            treeMap.put(A[i], i + 1); // 第i+1个数
        }

    }
}

