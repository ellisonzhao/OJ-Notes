package LeetCode.temp;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

class Solution {

    public int subarraysDivByK(int[] A, int K) {
        List<Integer> num = new ArrayList<Integer>();
        num.add(0);
        for (int i = 1; i < A.length; ++i) {
            num.add(A[i] + num.get(i - 1));
        }
        int ans = 0;
        for (int i = 0; i < num.size(); ++i) {
            for (int j = i + 1; j < num.size(); ++j) {
                if (((num.get(j) - num.get(i)) % K == 0))
                    ans++;
            }
        }
        return ans;
    }


    public List<String> generateParenthesis(int n) {
        List<String> ans = new ArrayList();
        backTrack(ans, "", 0, 0, n);
        return ans;
    }

    public void backTrack(List<String> ans, String curs, int lCount, int rCount, int count) {
        if (curs.length() == 2 * count) {
            ans.add(curs);
            return;
        }

        if (lCount < count) {
            backTrack(ans, curs + "(", lCount + 1, rCount, count);
        }

        if (rCount < lCount) {
            backTrack(ans, curs + ")", lCount, rCount + 1, count);
        }
    }

    public List<Integer> findDisappearedNumbers(int[] nums) {
        List<Integer> ans = new ArrayList<>();
        for (int i = 0; i < nums.length; ++i) {
            int val = Math.abs(nums[i]) - 1;
            if (nums[val] > 0) { //没有出现过的数(num)，则下标val=num-1对应的数不会被修改
                nums[val] = -nums[val];
            }
        }
        for (int i = 0; i < nums.length; ++i) {
            if (nums[i] > 0) {
                ans.add(i + 1);
            }
        }
        return ans;
    }


    public int totalFruit(int[] tree) {
        //求最长的连续子序列长度,最多包含两个不同元素
        int ans = 0, N = tree.length;
        Set<Integer> collected = new HashSet<>();
        // j用于遍历; i用于记录每种类型首次出现的index
        for (int i = 0, j = 0; j < N; ) {
            if (collected.size() <= 2) {
                collected.add(tree[j++]);
                //直接用下标计算长度;
//                ans = Math.max(ans, j - i);

            } else {
                ans = Math.max(ans, j - i - 1);
                collected.clear();
                for (int t = i; t < N; ++t) {
                    if (tree[t] != tree[i]) {
                        i = t;
                        j = i;
                        break;
                    }
                }
            }
            //一定要注意:循环结束之后,collected的容量可能大于2
            if (collected.size() <= 2) {
                ans = Math.max(ans, j - i);
            } else {
                ans = Math.max(ans, j - i - 1);
            }
        }
        return ans;
    }

    public int[] printMatrix(int[][] matrix) {
        if (matrix == null) return null;
        int rows = matrix.length, cols = matrix[0].length;
        int[] dx = {0, 1, 0, -1}, dy = {1, 0, -1, 0}, ans = new int[rows * cols]; //东、南、西、北
        boolean[][] visited = new boolean[rows][cols];
        int x = 0, y = 0, d = 0; //当前元素坐标，打印方向
        for (int i = 0; i < rows * cols; ++i) {
            ans[i] = matrix[x][y]; //x对应行数，y对应列数
            visited[x][y] = true;

            int r = x + dx[d], c = y + dy[d]; //当前位置沿着方向d的下一个元素
            if (c < 0 || c >= cols || r < 0 || r >= rows || visited[r][c]) {
                // 下一个元素是边界，则在当前位置改变方向
                d = (d + 1) % 4;
                // 改变方向后的下一个元素；
                r = x + dx[d];
                c = y + dy[d];
            }

            //当前位置更新为下一个元素
            x = r;
            y = c;
        }
        return ans;
    }

    public int binarySearch(int[] a, int fromIndex, int toIndex, int key) {
        int low = fromIndex;
        int high = toIndex;

        while (low <= high) {
            int mid = (low + high) >>> 1;
            int midVal = a[mid];

            if (midVal < key)
                low = mid + 1;
            else if (midVal > key)
                high = mid - 1;
            else
                return mid; // key found
        }
        return -(low + 1);  // key not found.
    }

    public int maxProduct(int[] A) {
        // store the result that is the max we have found so far
        int r = A[0], n = A.length;

        // imax/imin stores the max/min product of
        // subarray that ends with the current number A[i]
        for (int i = 1, imax = r, imin = r; i < n; i++) {
            // multiplied by a negative makes big number smaller, small number bigger
            // so we redefine the extremums by swapping them
            if (A[i] < 0) {
                int temp = imax;
                imax = imin;
                imin = temp;
            }

            // max/min product for the current number is either the current number itself
            // or the max/min by the previous number times the current one
            imax = Math.max(A[i], imax * A[i]);
            imin = Math.min(A[i], imin * A[i]);

            // the newly computed max value is a candidate for our global result
            r = Math.max(r, imax);
        }
        return r;
    }
}

public class Main {
    public static void main(String[] args) {
        int[] nums = {2, 3, -1, 2, 4};
        System.out.println(new Solution().maxProduct(nums));
    }
}