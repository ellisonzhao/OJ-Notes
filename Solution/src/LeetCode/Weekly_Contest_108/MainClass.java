package LeetCode.Weekly_Contest_108;

import java.util.HashSet;
import java.util.Set;

class Solution {
    public int numUniqueEmails(String[] emails) {
        Set<String> ret = new HashSet<>();
        for (String email : emails) {
            StringBuilder sb = new StringBuilder();
            int pos = email.indexOf('@');
            for (int i = 0; i < pos; ++i) {
                char ch = email.charAt(i);
                if (ch == '.')
                    continue;
                else if (ch == '+')
                    break;
                else
                    sb.append(ch);
            }
            sb.append(email.substring(pos));
            ret.add(sb.toString());
        }
        return ret.size();
    }

    public int numSubarraysWithSum(int[] A, int S) {
        int N = A.length, ans = 0;
        int[] sum = new int[N + 1], counter = new int[3000 + 5];
        counter[0] = 1;
        for (int i = 0; i < N; i++) {
            sum[i + 1] = sum[i] + A[i];
            counter[sum[i]]++;
        }

        for (int i = 0; i < N; ++i) {
            if (sum[i] - S >= 0) {
                ans += counter[sum[i] - S];
            }
        }
        return ans;
    }
}

public class MainClass {
    public static void main(String[] args) {
//        String[] emails = {"monthlyTest3_1.email+alex@leetcode.com", "monthlyTest3_1.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"};
//        System.out.println(new Solution().numUniqueEmails(emails));
        int[] A = {1, 0, 1, 1, 0, 1};
        int S = 2;
        System.out.println(new Solution().numSubarraysWithSum(A, S));
    }
}
