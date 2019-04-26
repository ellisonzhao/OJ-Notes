package AcWing.num_728;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.Queue;

/**
 * @author charlsonz
 * @date 2019-04-22 13:35
 */
public class Main {
    private static final int N = 10;
    static int[][] grid = new int[N][N];
    static int[][] time = new int[N][N];
    static int rows;
    static int cols;

    private static int bfs() {
        Queue<int[]> queue = new LinkedList<>();
        for (int[] t : time)
            Arrays.fill(t, -1);

        for (int i = 0; i < rows; i++) {
            for (int j = 0; j < cols; j++) {
                if (grid[i][j] == 2) {
                    time[i][j] = 0;
                    queue.add(new int[]{i, j});
                }
            }
        }

        int[] dx = new int[]{-1, 0, 1, 0}, dy = new int[]{0, 1, 0, -1};
        while (!queue.isEmpty()) {
            int[] t = queue.poll();
            int x = t[0]; // 行号
            int y = t[1]; // 列号
            int d = time[x][y]; // 从开始经历多少分钟
            for (int i = 0; i < 4; ++i) {
                int r = x + dx[i];
                int c = y + dy[i];
                if (r >= 0 && r < rows && c >= 0 && c < cols && time[r][c] == -1 && grid[r][c] == 1) {
                    time[r][c] = d + 1;
                    queue.add(new int[]{r, c});
                }
            }
        }

        int res = 0;
        for (int i = 0; i < rows; ++i) {
            for (int j = 0; j < cols; ++j) {
                if (grid[i][j] == 1) {
                    if (time[i][j] == -1) {
                        return -1;
                    }
                    res = Math.max(res, time[i][j]);
                }
            }
        }
        return res;
    }

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        String line;
        while (true) {
            line = br.readLine();
            if (line == null || line.length() == 0)
                break;

            String[] parts = line.trim().split(" ");
            int c = 0;
            for (int i = 0; i < parts.length; ++i) {
                grid[rows][c++] = Integer.parseInt(parts[i]);
            }
            cols = c;
            ++rows;
        }

        System.out.println(bfs());
    }
}
