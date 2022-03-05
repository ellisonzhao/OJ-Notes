class Solution {
 public:
  int maxCount;
  int count;
  vector<int> result;
  TreeNode* prev;
  vector<int> findMode(TreeNode* root) {
    count = 0;
    maxCount = 0;
    prev == nullptr;
    result.clear();
    search(root);

    return result;
  }
  void search(TreeNode* curr) {
    if (curr == nullptr) {
      return;
    }

    search(curr->left);

    if (prev == nullptr) {
      count = 1;
    } else if (prev->val == curr->val) {
      count++;
    } else {
      count = 1;
    }

    if (count == maxCount) {
      result.push_back(curr->val);
    }
    if (count > maxCount) {
      maxCount = count;
      result.clear();
      result.push_back(curr->val);
    }
    prev = curr;

    search(curr->right);

    return;
  }
};