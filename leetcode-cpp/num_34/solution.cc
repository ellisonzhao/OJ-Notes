class Solution {
 public:
  vector<int> searchRange(vector<int> &nums, int target) {
    if (nums.size() == 0) return {-1, -1};
    int first = searchFirst(nums, target);
    int last = searchLast(nums, target);
    return {first, last};
  }

  int searchFirst(vector<int> &nums, int target) {
    int l = 0, r = nums.size() - 1;
    while (l < r) {
      int mid = (l + r) >> 1;
      if (nums[mid] < target)
        l = mid + 1;
      else
        r = mid;
    }
    return nums[l] == target ? l : -1;
  }

  int searchLast(vector<int> &nums, int target) {
    int l = 0, r = nums.size() - 1;
    while (l < r) {
      int mid = (r + l + 1) >> 1;
      if (nums[mid] <= target)
        l = mid;
      else
        r = mid - 1;
    }
    return nums[l] == target ? l : -1;
  }
};