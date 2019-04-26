package AcWing.num_681.num_681_dfs;

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

    // 记录每个结点能到达的后代结点个数
    private static int[] state = new int[N];

    private static void dfs(int u) {
        state[u] = 1;
        for (int i = first[u]; i != -1; i = next[i]) {
            int v = edge[i];
            if (state[v] > 0)
                // 已经访问过
                continue;
            dfs(v);
            // 迭代形式避免重复计算
            state[u] += state[v];
        }
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

        dfs(1);
        int res = 0;
        // 对结点 1 的所有孩子结点进行 dfs
        for (int i = first[1]; i != -1; i = next[i]) {
            // i: 边的编号
            int u = edge[i];
            // 对从顶点1出发所有边对应的终点进行遍历
            res = Math.max(res, state[u]);
        }
        System.out.println(res);

    }
}