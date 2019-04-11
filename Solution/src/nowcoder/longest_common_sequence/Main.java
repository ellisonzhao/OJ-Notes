package nowcoder.longest_common_sequence;

public class Main {
    private static int findLongest(String A, int n, String B, int m) {
        // write code here
        if (A == null || B == null || n == 0 || m == 0)
            return 0;
        int[] dp = new int[n + 1];
        int res = 0;
        for (int i = 1; i <= m; i++) {
            for (int j = m; j >= 1; --j) {
                dp[j] = B.charAt(i - 1) == A.charAt(j - 1) ? dp[j - 1] + 1 : 0;
                res = Math.max(res, dp[j]);
            }
        }
        return res;
    }

    public int findLength(int[] A, int[] B) {
        if (A == null || B == null || A.length == 0 || B.length == 0) return 0;
        int m = A.length;
        int n = B.length;
        int res = 0;
        int[][] dp = new int[m + 1][n + 1];
        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                if (A[i - 1] == B[j - 1]) {
                    dp[i][j] = 1 + dp[i - 1][j - 1];
                    res = Math.max(res, dp[i][j]);
                }
            }
        }
        return res;
    }

    private static int findLongestCorrect(String A, int n, String B, int m) {
        // write code here
        char[] a = A.toCharArray();
        char[] b = B.toCharArray();
        int[][] dp = new int[n + 1][m + 1];
        int max = 0;
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= m; j++) {
                if (a[i - 1] == b[j - 1]) {
                    dp[i][j] = dp[i - 1][j - 1] + 1;
                    if (max < dp[i][j])
                        max = dp[i][j];
                }
            }
        }//for
        return max;
    }

    public static void main(String[] args) {
        String A = "cacccca", B = "aaacca";
        int n = 7, m = 6;
        System.out.println(findLongest(A, n, B, m));
        System.out.println(findLongestCorrect(A, n, B, m));
    }
}
