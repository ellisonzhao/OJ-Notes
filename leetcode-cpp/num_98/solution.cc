class Solution {
 public:
  TreeNode* prev;
  bool isValidBST(TreeNode* root) {
    if (root == nullptr) {
      return true;
    }

    bool isLeftValid = isValidBST(root->left);

    if (prev != nullptr && prev->val >= root->val) {
      return false;
    }
    prev = root;

    bool isRightValid = isValidBST(root->right);

    return isLeftValid && isRightValid;
  }
};