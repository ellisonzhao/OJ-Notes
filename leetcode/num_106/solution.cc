class Solution {
 public:
  unordered_map<int, int> inorderIndexes;
  TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder) {
    if (inorder.size() == 0 || postorder.size() == 0) {
      return nullptr;
    }
    for (size_t i = 0; i < inorder.size(); i++) {
      inorderIndexes[inorder[i]] = i;
    }
    return buildTree(inorder, 0, inorder.size() - 1, postorder, 0,
                     postorder.size() - 1);
  }
  TreeNode* buildTree(vector<int>& inorder, int inorderStart, int inorderEnd,
                      vector<int>& postorder, int postorderStart,
                      int postorderEnd) {
    if (inorderStart > inorderEnd || postorderStart > postorderEnd) {
      return nullptr;
    }

    int rootVal = postorder[postorderEnd];
    int leftLength = inorderIndexes[rootVal] - inorderStart;
    TreeNode* root = new TreeNode(rootVal);
    root->left =
        buildTree(inorder, inorderStart, inorderStart + leftLength - 1,
                  postorder, postorderStart, postorderStart + leftLength - 1);
    root->right =
        buildTree(inorder, inorderStart + leftLength + 1, inorderEnd, postorder,
                  postorderStart + leftLength, postorderEnd - 1);
    return root;
  }
};