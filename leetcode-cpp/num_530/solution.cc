class Solution {
 public:
  int result = INT_MAX;
  TreeNode* prev;
  void dfs(TreeNode* curr) {
    if (curr == nullptr) {
      return;
    }

    dfs(curr->left);

    if (prev != nullptr) {
      result = min(result, curr->val - prev->val);
    }
    prev = curr;

    dfs(curr->right);
  }
  int getMinimumDifference(TreeNode* root) {
    dfs(root);
    return result;
  }
};