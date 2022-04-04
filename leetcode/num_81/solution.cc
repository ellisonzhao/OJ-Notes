class Solution {
 public:
  bool search(vector<int> &nums, int target) {
    int k = nums.size() - 1;
    while (k > 0 && nums[0] == nums[k]) {
      k--;
    }

    int l = 0, r = k;
    while (l < r) {
      int mid = (l + r) / 2;
      if (nums[mid] >= nums[0])
        l = mid + 1;
      else
        r = mid;
    }

    return binary_search(nums.begin(), nums.begin() + l, target) ||
           binary_search(nums.begin() + l, nums.begin() + k + 1, target);
  }
};