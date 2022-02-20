/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left),
 * right(right) {}
 * };
 */
class Solution {
 public:
  int maxDepth = 0, bottomLeftValue;
  int findBottomLeftValue(TreeNode* root) {
    // 根结点深度为 1
    dfs(root, 1);
    return bottomLeftValue;
  }
  void dfs(TreeNode* root, int leftDepth) {
    // 传入 root 不能为空
    if (root == nullptr) {
      return;
    }
    // 正常递归退出条件
    if (root->left == nullptr && root->right == nullptr) {
      if (leftDepth > maxDepth) {
        maxDepth = leftDepth;
        bottomLeftValue = root->val;
      }
      return;
    }
    if (root->left) {
      // 递归到下一层，不修改当前 leftDepth 的值
      dfs(root->left, leftDepth + 1);
    }
    if (root->right) {
      // 递归到下一层，不修改当前 leftDepth 的值
      dfs(root->right, leftDepth + 1);
    }
  }
};