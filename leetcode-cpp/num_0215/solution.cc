//
// Created by ellisonzhao on 2022/9/24.
//

#include <bits/stdc++.h>

using namespace std;

class Solution {
 public:
  int partition(vector<int> &nums, int left, int right, int targetIndex) {
    // 随机生成枢纽元素
    int randomIndex = (rand() % (right - left + 1)) + left;
    swap(nums[left], nums[randomIndex]);
    int pivot = nums[left];
    // 目前枢纽元素在最左边，因此先从右往左找小于枢纽的元素
    int l = left, r = right;
    while (l < r) {
      // 从右往左找到第一个大于枢纽的元素
      while (l < r && nums[r] >= pivot) {
        --r;
      }
      nums[l] = nums[r];
      // 从左往右找到第一个小于枢纽的元素
      while (l < r && nums[l] <= pivot) {
        ++l;
      }
      nums[r] = nums[l];
    }
    nums[l] = pivot;

    if (l == targetIndex) {
      return nums[l];
    } else if (l > targetIndex) {
      return partition(nums, left, l - 1, targetIndex);
    } else {
      return partition(nums, l + 1, right, targetIndex);
    }
  }
  int findKthLargest(vector<int> &nums, int k) {
    // 初始化随机数种子
    srand((time(nullptr)));
    // 第 1 大的数，下标是 n - 1;
    // 第 2 大的数，下标是 n - 2;
    // ...
    // 第 k 大的数，下标是 n - k;
    int n = nums.size();

    return partition(nums, 0, n - 1, n - k);
  }
};