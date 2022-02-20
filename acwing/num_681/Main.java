package AcWing.num_681;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Arrays;

/**
 * @author charlsonz
 * @date 2019-04-20 09:11
 */
public class Main {

    private static final int N = 100010;

    private static int idx = 0; // 边的编号

    /**
     * first[u]=idx
     * 数组下标表示顶点 u
     * 值表示从顶点 u 出发的编号为idx的一条边
     * 并且是从顶点 u 出发最新添加到邻接表中的一条边
     */
    private static int[] first = new int[N]; //

    /**
     * next[idx] = first[u]
     * 数组下标表示编号为 idx 的边
     * 值表示编号为 idx 的边的“前一条边”的编号
     * 这里“前一条边”指的是first[u],
     * 也就是从顶点 u 出发的上一次添加到邻接表的边的编号,
     * 如果first[u]= -1, 表示当前边就是从顶点 u 出发的添加到邻接表的第一条边
     * first[u] != -1, 表示从顶点 u 出发已经有一条边添加到邻接表中
     */
    private static int[] next = new int[2 * N];

    /**
     * 添加(u,v)这条边, 顶点u开始，顶点v结束
     * 无向图的话需要添加(u,v)和(v,u)
     * edge[idx] = v
     * 表示编号为 idx 的这条边是以顶点 v 结束
     */
    private static int[] edge = new int[2 * N];

    private static void add(int u, int v) {
        // 添加当前编号为idx的边，并记录终点
        edge[idx] = v;

        // 记录这条边的前一条边
        next[idx] = first[u];

        // 更新first[u]为最新添加的这条边
        // 更新完后，边的编号加一
        first[u] = idx++;

    }

    private static int[] q = new int[N]; //数组表示队列,存放顶点

    // 标志数组, 值为true表示顶点已经访问
    private static boolean[] visited = new boolean[N];

    private static int bfs(int u) {
        q[0] = u;
        visited[u] = true;
        int front = 0; // 队头
        int rear = 0; // 队尾
        // 跟循环队列区别
        while (front <= rear) {
            int t = q[front++]; // 出队
            // 对顶点t 进行bfs
            for (int i = first[t]; i != -1; i = next[i]) {
                // 从最新添加的边一直遍历到 idx = -1
                int v = edge[i];
                if (!visited[v]) {
                    q[++rear] = v;
                    visited[v] = true;
                }
            }
        }
        return rear + 1;
    }

    private static int[] stringToIntegerArray(String input) {
        String[] parts = input.trim().split(" ");
        int[] output = new int[parts.length];
        for (int index = 0; index < parts.length; index++) {
            String part = parts[index].trim();
            output[index] = Integer.parseInt(part);
        }
        return output;
    }

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int n = Integer.parseInt(br.readLine());
        Arrays.fill(first, -1);

        while (--n > 0) {
            int[] parts = stringToIntegerArray(br.readLine());
            int u = parts[0];
            int v = parts[1];
            add(u, v);
            add(v, u);
        }

        int res = 0;
        visited[1] = true;
        for (int i = first[1]; i != -1; i = next[i]) {
            // 对从顶点1出发所有边对应的终点进行遍历
            res = Math.max(res, bfs(edge[i]));
        }
        System.out.println(res);

    }
}