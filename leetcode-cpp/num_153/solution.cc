class Solution {
 public:
  int findMin(vector<int> &nums) {
    int l = 0, r = nums.size();
    while (l < r) {
      int mid = l + (r - l) / 2;
      if (nums[mid] >= nums[0])
        l = mid + 1;
      else
        r = mid;
    }

    return l < nums.size() ? nums[l] : nums[0];
  }
};