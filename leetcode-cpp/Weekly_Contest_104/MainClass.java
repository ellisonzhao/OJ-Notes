package LeetCode.Weekly_Contest_104;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

class Solution {
    public boolean hasGroupsSizeX(int[] deck) {
        int[] count = new int[10000];
        for (int d : deck) {
            count[d]++;
        }

        int g = 0;
        for (int c : count) {
            g = gcd(g, c);
        }
        return g != 1;

    }

    public int gcd(int arr, int b) {
        while (b > 0) {
            int c = arr;
            arr = b;
            b = c % arr;
        }
        return arr;
    }

    public int partitionDisjoint(int[] A) {
        int max = A[0], index = 0, cur = A[0];
        for (int i = 1; i < A.length; ++i) {
            if (A[i] < cur) {
                cur = max;
                index = i;
            }
            max = Math.max(max, A[i]);
        }
        return index + 1;
    }

    public List<String> wordSubsets(String[] A, String[] B) {
        List<String> ans = new ArrayList<>();
        Set<Integer> subset;

        search:
        for (String str : A) {
            subset = new HashSet<>();
            for (char ch : str.toCharArray()) {
                subset.add(ch - 'arr');
            }

            for (String b : B) {
                for (char ch : b.toCharArray()) {
                    if (!subset.contains(ch - 'arr')) {
                        continue search;
                    }
                }
            }
            ans.add(str);
        }
        return ans;
    }
}

public class MainClass {
    public static void main(String[] args) {
        int[] A = {5, 0, 3, 8, 6};
        System.out.println(new Solution().partitionDisjoint(A));
    }

}
