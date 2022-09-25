//
// Created by ellisonzhao on 2022/9/25.
//

#include <bits/stdc++.h>

using namespace std;

class Solution {
 public:
  vector<vector<int>> threeSum(vector<int> &nums) {
    sort(nums.begin(), nums.end());
    vector<vector<int>> ans;
    for (int first = 0; first < nums.size() - 2; ++first) {
      // 后面的元素都大于 0，不可能找到结果
      if (nums[first] > 0) break;
      // 需要和上一次枚举的数不相同
      if (first > 0 && nums[first] == nums[first - 1]) {
        continue;
      }

      int second = first + 1, third = nums.size() - 1;
      int target = -nums[first];
      while (second < third) {
        if (nums[second] + nums[third] == target) {
          ans.push_back({nums[first], nums[second], nums[third]});
          // 跳过相同元素
          while (second < third && nums[second] == nums[second + 1]) ++second;
          while (second < third && nums[third] == nums[third - 1]) --third;
          ++second;
          --third;
        } else if (nums[second] + nums[third] < target) {
          ++second;
        } else {
          --third;
        }
      }
    }

    return ans;
  }
};