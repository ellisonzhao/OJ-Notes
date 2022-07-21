class Solution {
 public:
  vector<vector<int>> results;
  vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
    if (root == nullptr) {
      return results;
    }
    vector<int> path;
    dfs(root, path, targetSum);
    return results;
  }

  void dfs(TreeNode* node, vector<int> path, int sum) {
    if (node == nullptr) {
      return;
    }

    sum -= node->val;
    path.push_back(node->val);
    if (node->left == nullptr && node->right == nullptr && sum == 0) {
      results.push_back(path);
      return;
    }

    dfs(node->left, path, sum);
    dfs(node->right, path, sum);
  }
};