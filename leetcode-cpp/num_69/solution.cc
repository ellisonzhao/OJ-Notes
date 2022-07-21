class Solution {
 public:
  int mySqrt(int x) {
    int left = 0, right = x;
    while (left < right) {
      // 防止值溢出
      int mid = (right + left + 1LL) / 2;
      if (mid > x / mid)
        // 结果在左半区间 [left, mid-1]
        right = mid - 1;
      else
        // 结果在右半区间 [mid, right]
        left = mid;
    }
    return left;
  }
};