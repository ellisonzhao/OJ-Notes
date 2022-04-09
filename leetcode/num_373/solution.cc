class Solution {
 public:
  vector<vector<int>> kSmallestPairs(vector<int> &nums1, vector<int> &nums2,
                                     int k) {
    int n1 = nums1.size(), n2 = nums2.size();
    vector<vector<int>> ans;
    // 自定义排序，pair sum 小的在前，返回 false 使用大根堆
    auto cmp = [&](const auto &a, const auto &b) {
      return nums1[a.first] + nums2[a.second] >
             nums1[b.first] + nums2[b.second];
    };
    // 这里用小根堆，只需要找出最小的 k 个
    priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(cmp)> pq(
        cmp);
    // 开始多路归并
    for (int i = 0; i < min(n1, k); i++) {
      pq.push({i, 0});
    }
    // 出队当前 pair sum 最小的
    while (ans.size() < k && pq.size()) {
      auto [i, j] = pq.top();
      pq.pop();
      ans.push_back({nums1[i], nums2[j]});

      if (j + 1 < n2) pq.push({i, j + 1});
    }

    return ans;
  }
};