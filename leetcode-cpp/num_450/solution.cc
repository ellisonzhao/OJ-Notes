class Solution {
public:
  TreeNode *deleteNode(TreeNode *root, int key) {
    // 没找到待删除节点，遍历到空节点返回
    if (root == nullptr) {
      return root;
    }
    // 待删除节点在左子树
    if (root->val > key) {
      root->left = deleteNode(root->left, key);
      return root;
    }
    // 待删除节点在右子树
    if (root->val < key) {
      root->right = deleteNode(root->right, key);
      return root;
    }
    // 当前结点是待删除节点
    // 1. 左右孩子都为空（叶子节点），直接删除节点
    if (root->left == nullptr && root->right == nullptr) {
      // 内存释放！！！
      delete root;
      return nullptr;
    }
    // 2. 左孩子为空，右孩子不为空。删除当前节点只需用右孩子代替 ，返回右孩子
    if (root->left == nullptr) {
      TreeNode *node = root->right;
      // 内存释放！！！
      delete root;
      return node;
    }
    // 3. 左孩子不为空，右孩子为空，删除当前节点只需用左孩子代替 ，返回左孩子
    if (root->right == nullptr) {
      TreeNode *node = root->left;
      // 内存释放！！！
      delete root;
      return node;
    }
    // 5. 左右孩子均不为空，将当前根节点的左子树拼接到右子树最左节点的左孩子
    // 并返回删除节点右孩子为新的根节点。
    TreeNode *curr = root->right; // 找右子树最左节点
    while (curr->left != nullptr) {
      curr = curr->left;
    }
    // 把待删除节点（root）的左子树放在 curr 的左孩子位置
    curr->left = root->left;
    TreeNode *tmp = root; // 保存 root 节点，后续删除
    root = root->right;   // 返回旧 root 的右孩子作为新 root
    delete tmp;           // 释放节点内存

    return root;
  }
};