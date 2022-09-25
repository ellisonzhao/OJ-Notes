//
// Created by ellisonzhao on 2022/9/25.
//

#include <bits/stdc++.h>

using namespace std;

class Solution {
 public:
  int findKthLargest(vector<int> &nums, int k) {
    // 声明小顶堆
    priority_queue<int, vector<int>, greater<int>> pq;
    for (int num : nums) {
      pq.push(num);
      if (pq.size() > k) {
        pq.pop();
      }
    }
    return pq.top();
  }
};