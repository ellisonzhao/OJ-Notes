class Solution {
 public:
  TreeNode* trimBST(TreeNode* root, int low, int high) {
    if (root == nullptr) {
      return root;
    }
    // 说明左子树全部小于 low，只需要返回右子树的处理结果
    if (root->val < low) {
      return trimBST(root->right, low, high);
    }
    // 说明右子树全部大于 high，只需要返回左子树的处理结果
    if (root->val > high) {
      return trimBST(root->left, low, high);
    }
    // 当前节点不需要处理，递归处理左右子树
    root->left = trimBST(root->left, low, high);
    root->right = trimBST(root->right, low, high);
    return root;
  }
};