class Solution {
 public:
  double findMedianSortedArrays(vector<int> &nums1, vector<int> &nums2) {
    int n1 = nums1.size(), n2 = nums2.size();
    if (n1 > n2) return findMedianSortedArrays(nums2, nums1);

    int l = 0, r = n1;
    int k = (n1 + n2 + 1) / 2;  // 前 k 个 元素
    while (l < r) {
      int i = (l + r) / 2;
      int j = k - i;  // 保证划分之后，i + j = k
      // left_part 有 k 个元素，max(left_part) <= min(right_part)
      // 因为 B[j] >= B[j-1]，所以二分查找的条件就是找到 A[i] >= B[j-1] 的临界值
      //       left_part          |        right_part
      // A[0], A[1], ..., A[i-1]  |  A[i], A[i+1], ..., A[m-1]
      // B[0], B[1], ..., B[j-1]  |  B[j], B[j+1], ..., B[n-1]
      if (nums1[i] < nums2[j - 1])
        // A[i] 增加，B[j] 减小
        l = i + 1;
      else
        // A[i] 减小，B[j] 增加
        r = i;
    }

    int i = l, j = k - i;
    // 边界条件 i==0 => j==k，说明 A[0]>=B[j-1]，则返回 B[j-1]
    // 边界条件 j==0 => i==k， 说明 A[n1-1]<B[0]，则返回 A[i-1]
    int m1 =
        max(i == 0 ? INT_MIN : nums1[i - 1], j == 0 ? INT_MIN : nums2[j - 1]);
    // 如果两数组长度之和为奇数，那么 max(A[i-1], B[j-1])即为所求
    if ((n1 + n2) % 2 == 1) {
      return m1;
    }
    // 如果两数组长度之和为偶数，那么还要计算 min(A[i], B[j])
    // 边界条件 i==n1，说明 A[n1-1]<B[j-1]，则返回 B[j]
    // 边界条件 j==n2，说明 A[i]>=B[n2-1]，则返回 A[i]
    int m2 = min(i == n1 ? INT_MAX : nums1[i], j == n2 ? INT_MAX : nums2[j]);
    return (m1 + m2) / 2.0;
  }
};