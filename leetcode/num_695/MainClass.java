package LeetCode.num_695;

import java.util.Stack;

class Solution {
    public int maxAreaOfIsland(int[][] grid) {
        int ans = 0, ROW = grid.length, COL = grid[0].length;
        boolean[][] seen = new boolean[ROW][COL];

        int[] dx = new int[]{1, 0, -1, 0};
        int[] dy = new int[]{0, -1, 0, 1};

        for (int i = 0; i < ROW; ++i) {
            for (int j = 0; j < COL; ++j) {
                if ((grid[i][j] == 1) && (!seen[i][j])) {
                    int area = 0;
                    Stack<int[]> stack = new Stack<>();
                    stack.push(new int[]{i, j});
                    seen[i][j] = true;
                    while (!stack.isEmpty()) {
                        int[] temp = stack.pop();
                        int x = temp[0], y = temp[1];
                        ++area;
                        for (int k = 0; k < 4; ++k) {
                            int nx = x + dx[k];
                            int ny = y + dy[k];
                            if (nx >= 0 && nx < ROW &&
                                    ny >= 0 && ny < COL &&
                                    grid[nx][ny] == 1 &&
                                    !seen[nx][ny]) {
                                stack.push(new int[]{nx, ny});
                                seen[nx][ny] = true;

                            }
                        }
                    }
                    ans = Math.max(ans, area);
                }


            }
        }
        return ans;
    }

    public void moveZeroes(int[] nums) {
        int N = nums.length;
        for (int i = 0; i < N; ++i) {
            if (nums[i] == 0) {
                int j = i, temp;
                while (j < N) {
                    if (nums[j] != 0) {
                        temp = nums[j];
                        nums[j] = nums[i];
                        nums[i] = temp;
                        break;
                    }
                    ++j;
                }

            }
        }
    }
}

public class MainClass {
    public static void main(String[] args) {
        int[][] grids = new int[][]{
                {0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
                {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
                {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
                {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
                {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
                {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
                {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
                {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}
        };

        int[] nums = new int[]{0, 1, 0, 3,0,0, 12};
//        System.out.println(new Solution().maxAreaOfIsland(grids));
        new Solution().moveZeroes(nums);
        for (int num : nums)
            System.out.print(num+" ");
    }
}
