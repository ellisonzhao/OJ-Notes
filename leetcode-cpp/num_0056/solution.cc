//
// Created by ellisonzhao on 2023/2/12.
//

#include <bits/stdc++.h>

using namespace std;

class Solution {
 public:
  vector<vector<int>> merge(vector<vector<int>> &intervals) {
    sort(intervals.begin(), intervals.end());
    vector<vector<int>> merged;
    for (auto &interval : intervals) {
      int l = interval[0], r = interval[1];
      if (merged.empty() || merged.back()[1] < l) {
        // 与前一个区间没有交集
        merged.push_back({l, r});
      } else {
        // 要与前一个区间合并，更新区间右端点
        merged.back()[1] = max(merged.back()[1], r);
      }
    }
    return merged;
  }
};
