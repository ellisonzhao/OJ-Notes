package LeetCode.num_0200;

/**
 * @author ellison
 * @date 2022-10-22 11:33
 */
public class MainClass {

    class Solution {

        public int numIslands(char[][] grid) {
            int res = 0;
            for (int i = 0; i < grid.length; i++) {
                for (int j = 0; j < grid[0].length; j++) {
                    if (grid[i][j] == '1') {
                        dfsMarking(grid, i, j);
                        res++;
                    }
                }
            }
            return res;
        }

        private void dfsMarking(char[][] grid, int i, int j) {
            if (i >= 0 && j >= 0 && i < grid.length && j < grid[0].length && grid[i][j] == '1') {
                grid[i][j] = '0';
                dfsMarking(grid, i + 1, j);
                dfsMarking(grid, i - 1, j);
                dfsMarking(grid, i, j + 1);
                dfsMarking(grid, i, j - 1);
            }
        }
    }
}

class UnionFindSet {

    private int[] parents;
    private int[] ranks;
    int count;

    public UnionFindSet(char[][] grid, int rows, int cols) {
        parents = new int[rows * cols];
        ranks = new int[rows * cols];
        for (int i = 0; i < rows; ++i) {
            for (int j = 0; j < cols; ++j) {
                if (grid[i][j] == '1') {
                    int idx = i * cols + j;
                    parents[idx] = idx;
                    ++count;
                }
            }
        }
    }

    public int Find(int u) {
        if (parents[u] != u) {
            parents[u] = Find(parents[u]);
        }
        return parents[u];
    }

    public boolean Union(int u, int v) {
        int pu = Find(u);
        int pv = Find(v);

        if (pu == pv) {
            return false;
        }

        if (ranks[pu] > ranks[pv]) {
            parents[pv] = pu;
        } else if (ranks[pu] < ranks[pv]) {
            parents[pu] = pv;
        } else {
            parents[pv] = pu;
            ranks[pu] += 1;
        }
        --count;
        return true;
    }
}


class Solution {

    int[] dx = {-1, 0, 1, 0};
    int[] dy = {0, 1, 0, -1};

    public int numIslands(char[][] grid) {
        if (grid == null || grid.length == 0 || grid[0].length == 0) {
            return 0;
        }
        int rows = grid.length, cols = grid[0].length;
        UnionFindSet uf = new UnionFindSet(grid, rows, cols);
        for (int i = 0; i < rows; ++i) {
            for (int j = 0; j < cols; ++j) {
                if (grid[i][j] == '1') {
                    for (int dir = 0; dir < 4; ++dir) {
                        // 遍历到相邻位置
                        int x = i + dx[dir];
                        int y = j + dy[dir];
                        if (x >= 0 && x < rows && y >= 0 && y < cols && grid[x][y] == '1') {
                            int u = i * cols + j;
                            int v = x * cols + y;
                            uf.Union(u, v);
                        }
                    }
                }
            }
        }
        return uf.count;
    }
}
