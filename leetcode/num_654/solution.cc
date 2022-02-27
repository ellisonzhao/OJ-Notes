class Solution {
 public:
  TreeNode* constructMaximumBinaryTree(vector<int>& nums) {
    if (nums.size() == 0) {
      return nullptr;
    }

    return buildTree(nums, 0, nums.size() - 1);
  }
  TreeNode* buildTree(vector<int>& nums, int start, int end) {
    if (start > end) {
      return nullptr;
    }

    int maxVal = nums[start], idxOfMaxVal = start;
    // 注意区间是 [start, end]
    for (int i = start + 1; i <= end; i++) {
      if (nums[i] > maxVal) {
        maxVal = nums[i];
        idxOfMaxVal = i;
      }
    }

    TreeNode* root = new TreeNode(maxVal);
    root->left = buildTree(nums, start, idxOfMaxVal - 1);
    root->right = buildTree(nums, idxOfMaxVal + 1, end);

    return root;
  }
};