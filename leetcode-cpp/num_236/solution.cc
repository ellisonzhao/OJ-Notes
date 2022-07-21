class Solution {
 public:
  TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
    if (root == nullptr || root == p || root == q) {
      return root;
    }

    TreeNode* left = lowestCommonAncestor(root->left, p, q);
    TreeNode* right = lowestCommonAncestor(root->right, p, q);
    // p 在左子树、q 在右子树；
    // 或者 p 在右子树、q 在左子树
    if (left != nullptr && right != nullptr) {
      return root;
    }
    if (left == nullptr) {
      return right;
    }

    return left;
  }
};