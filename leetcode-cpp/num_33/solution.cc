class Solution {
 public:
  int search(vector<int> &nums, int target) {
    if (nums.empty()) return -1;

    int pos = searchPivot(nums, 0, nums.size() - 1, nums[0]);
    int leftIndex = searchPivot(nums, 0, pos - 1, target);
    if (nums[leftIndex] == target) return leftIndex;

    int rightIndex = searchPivot(nums, pos, nums.size() - 1, target);
    if (nums[rightIndex] == target) return rightIndex;

    return -1;
  }

  // [left, right] 二分查找旋转数组分界点
  int searchPivot(vector<int> &nums, int left, int right, int target) {
    while (left < right) {
      int mid = (left + right) / 2;
      if (nums[mid] >= target)
        left = mid + 1;
      else
        right = mid;
    }
    return left;
  }

  // [left, right] 查找target
  int binarySearch(vector<int> &nums, int left, int right, int target) {
    while (left < right) {
      int mid = (left + right) / 2;
      if (nums[mid] < target)
        left = mid + 1;
      else
        right = mid;
    }
    return left;
  }
};
