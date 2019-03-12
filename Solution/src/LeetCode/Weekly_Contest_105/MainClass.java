package LeetCode.Weekly_Contest_105;

import java.util.ArrayDeque;
import java.util.Arrays;
import java.util.Deque;

class Solution {

    public int maxSubarraySumCircular(int[] A) {
        int N = A.length;

        // Compute P[j] = B[0] + B[1] + ... + B[j-1]
        // for fixed array B = A+A
        int[] P = new int[2 * N + 1];
        for (int i = 0; i < 2 * N; ++i)
            P[i + 1] = P[i] + A[i % N];

        // Want largest P[j] - P[i] with 1 <= j-i <= N
        // For each j, want smallest P[i] with i >= j-N
        int ans = A[0];
        // deque: i's, increasing by P[i]
        Deque<Integer> deque = new ArrayDeque();
        deque.offer(0);

        for (int j = 1; j <= 2 * N; ++j) {
            // If the smallest i is too small, remove it.
            if (deque.peekFirst() < j - N)
                deque.pollFirst();

            // The optimal i is deque[0], for cand. answer P[j] - P[i].
            // B[i] + B[i+1] + ... + B[j-1]
            ans = Math.max(ans, P[j] - P[deque.peekFirst()]);

            // Remove any i1's with P[i2] <= P[i1].
            /**
             *  关键一步,类似于dp[j+1]=A[j+1]+max(dp[j],0)
             *  p[j]
             *  更新队列:
             *
             */
            while (!deque.isEmpty() && P[j] <= P[deque.peekLast()])
                deque.pollLast();

            deque.offerLast(j);
        }

        return ans;
    }

    public int smallestRangeI(int[] A, int K) {
        int N = A.length, min = A[0], max = A[0];
        for (int i = 0; i < N; ++i) {
            min = Math.min(min, A[i]);
            max = Math.max(max, A[i]);
        }
        return (K >= (max - min) / 2) ? 0 : (max - min - 2 * K);
    }

    public int smallestRangeII(int[] A, int K) {
        Arrays.sort(A);
        int N = A.length, ans = A[N - 1] - A[0];
        for (int i = 0; i < N - 1; ++i) {
            int max = Math.max(A[i] + K, A[N - 1] - K);
            int min = Math.min(A[0] + K, A[i + 1] - K);
            ans = Math.min(ans, max - min);
        }
        return ans;
    }

    public int minAddToMakeValid(String S) {
        int ans = 0, bracket = 0;
        for (char c : S.toCharArray()) {
            if (c == '(') {
                ++bracket;
            } else {
                if (bracket > 0) {
                    --bracket;
                } else {
                    ++ans;
                }
            }
        }
        return ans + bracket;
    }

    public int threeSumMulti(int[] A, int target) {
        long ans = 0, MOD = 1_000_000_000 + 7;
        int N = A.length;
        Arrays.sort(A);

        for (int i = 0; i < N; ++i) {
            int T = target - A[i];
            // i < j < k
            int j = i + 1, k = N - 1;
            while (j < k) {
                if (A[j] + A[k] < T)
                    j++;
                else if (A[j] + A[k] > T)
                    k--;
                else {  // A[j] + A[k] == T
                    if (A[j] != A[k]) {
                        //需要分别统计A[j] 和 A[k]的个数,然后计算组合数
                        int count_j = 1, count_k = 1;
                        while (j + 1 < k && A[j] == A[j + 1]) {
                            count_j++;
                            j++;
                        }
                        while (k - 1 > j && A[k] == A[k - 1]) {
                            count_k++;
                            k--;
                        }
                        ans += count_j * count_k;
                        j++;
                        k--;
                    } else { // A[j] == A[k]
                        // count_j==count_k==k-j+1;
                        ans += (k - j + 1) * (k - j) / 2;
                        break;
                    }
                }
            }
        }


        return (int) (ans % MOD);
    }
}

public class MainClass {
    public static void main(String[] args) {
        int[] A = {1, 1, 2, 2, 3, 3, 4, 4, 5, 5};
        int target = 8;
        System.out.println(new Solution().threeSumMulti(A, target));
    }
}
