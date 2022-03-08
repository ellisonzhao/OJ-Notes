class Solution {
 public:
  TreeNode* buildNode(vector<int>& nums, int left, int right) {
    if (left > right) return nullptr;

    int mid = left + ((right - left) / 2);
    TreeNode* root = new TreeNode(nums[mid]);
    root->left = buildNode(nums, left, mid - 1);
    root->right = buildNode(nums, mid + 1, right);

    return root;
  }
  TreeNode* sortedArrayToBST(vector<int>& nums) {
    return buildNode(nums, 0, nums.size() - 1);
  }
};