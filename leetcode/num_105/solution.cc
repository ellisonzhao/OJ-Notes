class Solution {
 public:
  // 保存中序序列各个节点的索引位置，便于计算左子树序列长度（节点个数）
  unordered_map<int, int> inorderIndexes;

  TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
    if (preorder.size() == 0 || inorder.size() == 0) {
      return nullptr;
    }
    for (size_t i = 0; i < inorder.size(); i++) {
      inorderIndexes[inorder[i]] = i;
    }
    return buildTree(preorder, 0, preorder.size() - 1, inorder, 0,
                     inorder.size() - 1);
  }

  TreeNode* buildTree(vector<int>& preorder, int preStart, int preEnd,
                      vector<int>& inorder, int inStart, int inEnd) {
    if (preStart > preEnd || inStart > inEnd) {
      return nullptr;
    }

    int rootVal = preorder[preStart];
    int leftLength = inorderIndexes[rootVal] - inStart;
    TreeNode* root = new TreeNode(rootVal);
    root->left = buildTree(preorder, preStart + 1, preStart + leftLength,
                           inorder, inStart, inStart + leftLength - 1);
    root->right = buildTree(preorder, preStart + leftLength + 1, preEnd,
                            inorder, inStart + leftLength + 1, inEnd);
    return root;
  }
};