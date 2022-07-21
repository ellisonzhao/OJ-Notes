class Solution {
 public:
  int sum;
  void addSum(TreeNode* root) {
    if (root == nullptr) return;

    addSum(root->right);
    sum += root->val;
    root->val = sum;
    addSum(root->left);
  }
  TreeNode* convertBST(TreeNode* root) {
    if (root == nullptr) return root;

    sum = 0;
    addSum(root);
    return root;
  }
};