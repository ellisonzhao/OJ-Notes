//
// Created by ellisonzhao on 2022/10/22.
//

#include <bits/stdc++.h>

using namespace std;

const int MAX_N = 100000;
class Solution {
 public:
  int count;
  int parents[MAX_N];

  int find(int i) {
    if (parents[i] == i) return parents[i];
    parents[i] = find(parents[i]);
    return parents[i];
  }

  void unite(int i, int j) {
    int u = find(i);
    int v = find(j);
    if (u == v) return; // 避免重复合并操作
    parents[u] = parents[v];
    count--;
  }

  int numIslands(vector<vector<char>> &grid) {
    int m = grid.size();
    int n = grid[0].size();

    // 初始化 parent 数组，记录初始岛屿数（也就是 1 的数目）
    for (int i = 0; i < m; i++) {
      for (int j = 0; j < n; j++) {
        int idx = i * n + j;
        parents[idx] = idx;
        if (grid[i][j] == '1')
          count++;
      }
    }

    for (int i = 0; i < m; i++) {
      for (int j = 0; j < n; j++) {
        if (grid[i][j] == '0') continue;

        int idx = i * n + j;
        if (i + 1 < m && grid[i + 1][j] == '1') { //合并岛屿
          unite(idx, (i + 1) * n + j);
        }
        if (j + 1 < n && grid[i][j + 1] == '1') {
          unite(idx, i * n + j + 1);
        }

      }
    }

    return count;
  }
};